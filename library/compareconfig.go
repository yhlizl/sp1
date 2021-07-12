package library

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gogf/gf/frame/g"
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

//CompareConfig is Copare path location all config and point out difference
func CompareConfig(root string) g.Map {
	datalist := g.Map{}
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		//path 包含檔名

		datalist[info.Name()] = ReadConfig(path)
		fmt.Printf("%v/n", datalist[info.Name()])
		return nil
	})
	return datalist
}
