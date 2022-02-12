package config

import (
	"encoding/json"
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
		log.Fatal(err)
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

func Save(newConfig Config) {
	currentConfig = newConfig
	file, err := json.MarshalIndent(currentConfig, "", "  ")
	err = ioutil.WriteFile("config.json", file, 0644)
	if err != nil {
		log.Println(err)
	}

}

var currentConfig Config

func UpdateFrequency(value float64) {
	newConfig := Get()
	newConfig.Frequency = value
	Save(newConfig)
}

func GetFrequency() float64 {
	return Get().Frequency
}
func UpdateMultiplier(value float64) {
	newConfig := Get()
	newConfig.Multiplier = value
	Save(newConfig)
}

func GetMultiplier() float64 {
	return Get().Multiplier
}

func GetPort() uint16 {
	return Get().Port
}

func GetYouTubeAPIKey() string {
	return Get().YouTubeAPIKey
}

func GetSSD1306() bool {
	return Get().SSD1306
}
