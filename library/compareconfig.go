package library

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

//ReadConfig is getiing config hash map and reuturn
func ReadConfig(path string) map[string]interface{} {
	v := viper.New()
	// viper.SetConfigName("config")         // name of config file (without extension)
	v.SetConfigType("properties") // REQUIRED if the config file does not have the extension in the name
	// json
	// yaml
	// ini
	// toml
	// properties
	// viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("not found file!")
		} else {
			// Config file was found but another error was produced
			log.Println("found file, but error !")

		}
	}
	return v.AllSettings()
}

//CompareConfig is Copare path location all config and point out difference,return map data and diff map, map[params][tools]=value
func CompareConfig(root string) (map[string]map[string]interface{}, map[string]bool) {
	datalist := make(map[string]map[string]interface{})
	count := 0
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		//path 包含檔名
		count++
		datalist[info.Name()] = ReadConfig(path)
		fmt.Printf("%v/n", datalist[info.Name()])
		return nil
	})
	configMap := make(map[string]map[string]interface{})
	temp := make(map[string]interface{})

	//get all config into mapmap
	for configName, c := range datalist {
		for iName, iValue := range c {
			temp[configName] = iValue
			if _, ok := configMap[iName]; !ok {
				configMap[iName] = temp
			} else {
				configMap[iName][configName] = iValue
			}
		}

	}

	//dig out difference
	difName := map[string]bool{}
	for iName, c := range configMap {
		temp := ""
		for _, value := range c {
			if temp == "" {
				temp = value.(string)
				continue
			}
			if value != temp {
				difName[iName] = true
			}
		}
		if len(c) != count {
			difName[iName] = true
		}
	}

	return configMap, difName
}
