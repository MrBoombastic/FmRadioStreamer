package config

import (
	"encoding/json"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"os"
)

// Get returns current config saved in config file with mutex.
func Get() (*SafeConfig, error) {
	data, err := os.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}
	cfg := new(SafeConfig)
	err = json.Unmarshal(data, &cfg.Config)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Save saves new config to config file. Needs mutex locked!
func Save(cfg *Config) {
	file, err := json.MarshalIndent(cfg, "", "  ")
	err = os.WriteFile("config.json", file, 0644)
	if err != nil {
		logs.FmRadStrFatal(err)
	}
}
