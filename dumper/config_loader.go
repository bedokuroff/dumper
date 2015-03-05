package dumper

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Dbstrings []string `json:"dbstrings"`
	WriteStep uint32   `json:"writeStep"`
	ReadStep  uint32   `json:"readStep"`
}

func LoadConfig(filename string) Configuration {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("There was an error loading configuration: ", err)
	}
	return config
}
