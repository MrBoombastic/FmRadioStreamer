package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Get() Config {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(err)
	}
	return configuration
}
