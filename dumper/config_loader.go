package dumper

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Dbstrings []string `json:"dbstrings"`
	WriteStep uint32   `json:"writeStep"`
	ReadStep  uint32   `json:"readStep"`
}

func LoadConfig(filename string) Configuration {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening the file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("There was an error loading configuration: ", err)
	}

	return config
}
