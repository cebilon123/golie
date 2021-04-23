package file

import (
	"os"
)

//GetConfigDirectory returns application directory which is placed inside user's AppData/roaming directory.
func GetConfigDirectory() string {
	appdataConfig, err := os.UserConfigDir()
	if err != nil {
		panic("There is no config directory")
	}

	return appdataConfig + "\\" + "golie"
}
