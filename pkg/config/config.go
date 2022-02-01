package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Get() Config {
	if (Config{}) != currentConfig { //If not empty
		return currentConfig
	}
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	configuration := Config{}
	if err := json.NewDecoder(file).Decode(&configuration); err != nil {
		log.Fatal(err)
	}
	return configuration
}

func GetMap() map[string]interface{} {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var data map[string]interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatal(err)
	}
	return data
}

func save(newConfig Config) {
	currentConfig = newConfig
	file, _ := json.MarshalIndent(currentConfig, "", "  ")
	_ = ioutil.WriteFile("config.json", file, 0644)
}

var currentConfig Config

func UpdateFrequency(value float64) {
	newConfig := Get()
	newConfig.Frequency = value
	save(newConfig)
}

func GetFrequency() float64 {
	return Get().Frequency
}
func UpdateMultiplier(value float64) {
	newConfig := Get()
	newConfig.Multiplier = value
	save(newConfig)
}

func GetMultiplier() float64 {
	return Get().Multiplier
}

func GetPort() uint16 {
	return Get().Port
}
