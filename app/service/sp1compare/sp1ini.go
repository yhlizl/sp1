package sp1compare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"sp1/app/service/ftp"
	"sp1/library"
	"strconv"
	"strings"
	"sync"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
)

var inipath string
var ininame string
var filepath string
var configpath string
var ftpmap map[string]map[string]map[string]map[string]interface{}
var syn sync.WaitGroup
var mu sync.Mutex

//ConfigType is struct for config
type ConfigType struct {
	Inipath    string
	Ininame    string
	Filepath   string
	Configpath string
}

//Config is initial
var Config ConfigType

//Configdata is for parse json to struct
type Configdata struct {
	Check  bool              `json:"check"`
	Params string            `json:"params"`
	Value  map[string]string `json:"value"`
}

//Configdatas is for parse json to struct
type Configdatas struct {
	Configdata []Configdata `json:"configdatas"`
}

//TODO: modify cfg
//Init is initialize
func init() {
	g.Cfg().AddPath("/Users/royale/go/src/sp1/config/")

	// ftpcfg := g.Cfg().GetMap("ftp")
	// ftpmap = make(map[string]map[string]interface{})
	// for i, v := range ftpcfg {
	// 	ftpmap[i] = v.(map[string]interface{})
	// }
	ftpcfg := g.Cfg().GetMap("ftp")
	ftpmap = make(map[string]map[string]map[string]map[string]interface{})
	for i, v := range ftpcfg {
		ftpmap[i] = map[string]map[string]map[string]interface{}{}
		temp := v.([]interface{})
		for in, v2 := range temp {
			ind := strconv.Itoa(in)
			ftpmap[i][ind] = map[string]map[string]interface{}{}
			temp2 := v2.(map[string]interface{})
			for in3, v3 := range temp2 {
				ftpmap[i][ind][in3] = v3.(map[string]interface{})
			}
		}
	}

	inipath = g.Cfg().GetString("sp1ini.root")
	ininame = g.Cfg().GetString("sp1ini.filename")
	filepath = g.Cfg().GetString("filesystem.path")
	configpath = g.Cfg().GetString("sp1ini.configpath")
	Config = ConfigType{inipath, ininame, filepath, configpath}
}
func inSlice(list []string, str string) bool {
	for _, lists := range list {
		//fmt.Println("start compare:", lists, str)
		if strings.Compare(lists, str) == 0 {
			//	fmt.Println("same compare:", lists, str)

			return true
		}
	}
	return false
}
func downloadConfig(fab, phase []string) []string {
	log.Println("FTP Config get ;", ftpmap)
	filelist := []string{}
	params := ""
	for _, f := range fab {
		params += f
	}
	for _, p := range phase {
		params += p
	}

	ch := make(chan bool, 10)
	for i, v := range ftpmap {
		if !(inSlice(fab, i)) {
			//fmt.Println("gaptss fab :", i, fab)
			continue
		}
		for _, phs := range v {
			for pha, v2 := range phs {
				if !(inSlice(phase, pha)) {
					//	fmt.Println("gaptss phase :", pha)

					continue
				}
				syn.Add(1)
				ch <- true
				go func(v2 map[string]interface{}, filelist *[]string) {
					c := ftp.ConnnectFTP(v2["url"].(string), v2["port"].(string), v2["user"].(string), v2["pwd"].(string))
					file := ftp.GetFromFTP(c, inipath, filepath, ininame, v2["name"].(string), params)
					mu.Lock()
					*filelist = append(*filelist, file...)
					mu.Unlock()
					fmt.Println("start to download : ", inipath, filepath, ininame, v2["name"].(string), params)
					if err := c.Quit(); err != nil {
						log.Fatal(err)
					}

					<-ch
					syn.Done()

				}(v2, &filelist)

			}

		}
	}
	syn.Wait()
	return filelist
}

//Compareini is downloading from ftp and comprare and write config to config.ini
func Compareini(fab, phase []string) []interface{} {

	res := g.Array{}
	filelist := downloadConfig(fab, phase)
	dataftpmap, diff := library.CompareConfigList(filepath, filelist)
	params := ""
	for _, f := range fab {
		params += f
	}
	for _, p := range phase {
		params += p
	}

	for i, v := range dataftpmap {
		temp := gmap.New()
		temp.Set("params", i)
		temp.Set("check", diff[i])
		result := map[string]string{}

		for ind, val := range v {

			result[ind] = val.(string)

		}
		temp.Set("value", result)
		res = append(res, temp)

	}
	WriteToConfig(res, params)
	return res
}

//WriteToConfig is to write config to config/sp1/ini.config
func WriteToConfig(data []interface{}, params string) string {
	log.Println("=====Start to write sp1 config ====")

	path := path.Join(configpath, params)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Println(err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.Write([]byte("["))
	if err != nil {
		log.Println(err)
	}
	for i, v := range data {
		u, err := json.Marshal(v)
		if err != nil {
			log.Println(err)
		}
		_, err = buf.Write(u)
		if err != nil {
			log.Println(err)
		}
		if i != len(data)-1 {
			_, err = buf.Write([]byte(","))
			if err != nil {
				log.Println(err)
			}
		}
	}
	_, err = buf.Write([]byte("]"))
	if err != nil {
		log.Println(err)
	}
	_, err = buf.WriteTo(file)
	if err != nil {
		log.Println(err)
	}
	return path
}

//ReadFromConfig is to read config to config/sp1/ini.config
func ReadFromConfig(path string) []Configdata {
	log.Println("=====Start to read sp1 config ====")
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		log.Println(err)
	}
	res := buf.String()
	fmt.Println(res)

	// resjson, err := json.Marshal(res)
	// if err != nil {
	// 	log.Println(err)
	// }
	var d []Configdata
	err = json.Unmarshal([]byte(res), &d)
	if err != nil {
		log.Println(err)
	}
	log.Println("==================")
	log.Println("Get config : ", d)
	return d
}
