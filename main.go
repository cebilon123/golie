package main

import (
	"github.com/cebilon123/golie/config"
	"github.com/cebilon123/golie/file"
)

var initConfig = &config.Init{}

func main() {
	//TODO Introduce CLI lib to create commands
	initConfig = config.LoadInitConfig(file.GetConfigDirectory())
}
