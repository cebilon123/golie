package cfg

import (
	"encoding/binary"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"os"
)

//Configuration represents serialized configuration for user.
type Configuration struct {
	Path string `json:"path"`
}

//Equal function to check if configurations are the same.
func (c *Configuration) Equal(c2 *Configuration) bool{
	return c.Path == c2.Path
}

//TODO Write tests for this OverwriteWithDifferentFields()

//OverwriteWithDifferentFields overwrites configuration with different values of c2.
func (c *Configuration) OverwriteWithDifferentFields(c2 *Configuration) {
	//struct to overwrite
	m := structs.Map(c)
	m2 := structs.Map(c2)

	for key := range m {
		if m2[key] != nil && m[key] != m2[key] {
			m[key] = m2[key]
		}
	}

	if err := mapstructure.Decode(m, &c); err != nil {
		println(err)
		return
	}
}

const (
	configurationFileName = "config.bin"
)

func Deserialize() *Configuration {

}

//Serialize serializes configuration into file. If there is existing config file it checks differences and overwrites only
//different fields.
func (c *Configuration) Serialize() {
	existingConfig := Deserialize()
	if existingConfig != nil && !c.Equal(existingConfig){
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
