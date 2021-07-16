package sp1compare

import (
	"log"
	"sp1/app/service/ftp"
	"sp1/library"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
)

var inipath string
var ininame string
var filepath string
var ftpmap map[string]map[string]interface{}

//TODO: modify cfg
//Init is initialize
func init() {
	g.Cfg().AddPath("/Users/royale/go/src/sp1/config/")

	ftpcfg := g.Cfg().GetMap("ftp")
	ftpmap = make(map[string]map[string]interface{})
	for i, v := range ftpcfg {
		ftpmap[i] = v.(map[string]interface{})
	}
	inipath = g.Cfg().GetString("sp1ini.root")
	ininame = g.Cfg().GetString("sp1ini.filename")
	filepath = g.Cfg().GetString("filesystem.path")

}

func downloadConfig() {
	log.Println("FTP Config get ;", ftpmap)
	for i, v := range ftpmap {
		c := ftp.ConnnectFTP(v["url"].(string), v["port"].(string), v["user"].(string), v["pwd"].(string))
		ftp.GetFromFTP(c, inipath, filepath, ininame, i)

		if err := c.Quit(); err != nil {
			log.Fatal(err)
		}
	}
}

//Compareini is downloading from ftp and comprare
func Compareini() []interface{} {
	data, diff := library.CompareConfig(filepath)

	res := g.Array{}
	downloadConfig()

	for i, v := range data {
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

	return res
}
