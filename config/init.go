package config

import (
	"encoding/json"
	"fmt"
	"os"
	path2 "path"
)

//Init is an initial config.
type Init struct {
	//Path is a location of directory in which all files are watched and synced with directory on other machine.
	Path string `json:"path"`
}

const configFileName = "golie.config"

//LoadInitConfig loads initial config file which is json file type.
func LoadInitConfig(directoryPath string) *Init {
	path := path2.Join(directoryPath, configFileName)

	content, err := os.ReadFile(path)

	if err != nil {
		//config file probably do not exists
		if !createInitConfig(path) {
			panic("Cannot create golie.config")
		}
	}

	init := &Init{}
	if err := json.Unmarshal(content, init); err != nil {
		panic(fmt.Sprint("Error while deserializing config: ", err))
	}

	return init
}

//createInitConfig creates initial config if it not exists.
func createInitConfig(path string) bool {
	if _, err := os.Create(path); err != nil {
		return false
	}

	//TODO force user to create configuration for path

	return true
}
