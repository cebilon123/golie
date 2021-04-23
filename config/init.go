package config

import (
	"os"
)

//Init is an initial config.
type Init struct {
	//Path is a location of directory in which all files are watched and synced with directory on other machine.
	Path string `json:"path"`
}

const (
	configFileName = "golie.config"
)

//ReadInitConfig reads initial config into binary slice which should be deserialized into Init struct.
func ReadInitConfig(directoryPath string) []byte {
	path := directoryPath + "\\" + configFileName

	content, err := os.ReadFile(path)

	if err != nil {
		//config file probably do not exists
		if !createInitConfig(directoryPath, path) {
			panic("Cannot create golie.config")
		}
	}

	return content
}

//createInitConfig creates initial config if it not exists.
func createInitConfig(directoryPath string, path string) bool {
	if err := os.MkdirAll(directoryPath, os.ModePerm); err != nil {
		return false
	}

	if _, err := os.Create(path); err != nil {
		return false
	}

	//TODO force user to create configuration for path

	return true
}
