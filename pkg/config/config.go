package config

import (
	"encoding/json"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"os"
)

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

// Save saves new config. Needs mutex locked!
func Save(cfg *SafeConfig) {
	file, err := json.MarshalIndent(cfg.Config, "", "  ")
	err = os.WriteFile("config.json", file, 0644)
	if err != nil {
		logs.PiFmAdvError(err)
	}
}
