package cfg

import (
	"encoding/binary"
	"os"
)

const (
	configurationFileName = "config.bin"
)

//Configuration represents serialized configuration for user.
type Configuration struct {
	Path string `json:"path"`
}

//Equal function to check if configurations are the same.
func (c *Configuration) Equal(c2 *Configuration) bool {
	return c.Path == c2.Path
}

//OverwriteWithDifferentFields overwrites configuration with different values of c2. This is simplified (with use of ifs)
//in sake of speed and efficiency.
func (c *Configuration) OverwriteWithDifferentFields(c2 *Configuration) {
	if (c.Path != c2.Path || len(c.Path) == 0) && len(c2.Path) > 0 {
		c.Path = c2.Path
	}
}

func Deserialize() *Configuration {
	return nil
}

//Serialize serializes configuration into file. If there is existing config file it checks differences and overwrites only
//different fields.
func (c *Configuration) Serialize() {
	existingConfig := Deserialize()
	if existingConfig != nil && !c.Equal(existingConfig) {
		c.OverwriteWithDifferentFields(existingConfig)
	}

	f, err := os.Create(configurationFileName)
	if err != nil {
		println(err)
	}

	err = binary.Write(f, binary.LittleEndian, c)
	if err != nil {
		println(err)
	}
}
