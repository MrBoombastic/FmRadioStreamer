package config

import (
	"encoding/json"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"log"
	"os"
)

var currentConfig Config

// Get returns currently saved config
func Get() (Config, error) {
	if (Config{}) != currentConfig { //If not empty, return config from memory
		return currentConfig, nil
	}
	file, err := os.Open("config.json")
	if err != nil {
		return Config{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logs.FmRadStrFatal(err)
		}
	}(file)
	configuration := Config{}
	if err := json.NewDecoder(file).Decode(&configuration); err != nil {
		log.Fatal(err)
	}
	return configuration, nil
}

// GetMap returns currently saved config, but in map
func GetMap() map[string]interface{} {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	var data map[string]interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatal(err)
	}
	return data
}

// Save saves new config
func Save(newConfig Config) {
	currentConfig = newConfig
	file, err := json.MarshalIndent(currentConfig, "", "  ")
	err = os.WriteFile("config.json", file, 0644)
	if err != nil {
		logs.PiFmAdvError(err)
	}

}

// UpdateFrequency saves new frequency to config
func UpdateFrequency(value float64) {
	newConfig, _ := Get()
	newConfig.Frequency = value
	Save(newConfig)
}

// GetFrequency returns current frequency (Get wrapper)
func GetFrequency() float64 {
	config, _ := Get()
	return config.Frequency
}

// UpdateMultiplier saves new multiplier to config
func UpdateMultiplier(value float64) {
	newConfig, _ := Get()
	newConfig.Multiplier = value
	Save(newConfig)
}

// GetMultiplier returns current multiplier (Get wrapper)
func GetMultiplier() float64 {
	config, _ := Get()
	return config.Multiplier
}

// GetPort returns current dashboard port (Get wrapper)
func GetPort() uint16 {
	config, _ := Get()
	return config.Port
}

// GetYouTubeAPIKey returns current YT API (Get wrapper)
func GetYouTubeAPIKey() string {
	config, _ := Get()
	return config.YouTubeAPIKey
}

// GetSSD1306 returns current screen state in boolean (Get wrapper)
func GetSSD1306() bool {
	config, _ := Get()
	return config.SSD1306
}

// GetDynamicRTInterval returns current dynamic RT switching interval (Get wrapper)
func GetDynamicRTInterval() uint {
	config, _ := Get()
	return config.DynamicRTInterval
}

// GetVerbose returns if app should be spitting out stdout from child processes
func GetVerbose() bool {
	config, _ := Get()
	return config.Verbose
}

// GetFormat returns stored audio format
func GetFormat() string {
	config, _ := Get()
	return config.Format
}
