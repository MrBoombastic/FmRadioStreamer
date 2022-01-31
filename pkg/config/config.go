package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(err)
	}
	return configuration
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
