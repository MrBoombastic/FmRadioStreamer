package buttons

import (
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/ssd1306"
	"github.com/stianeikeland/go-rpio/v4"
)

var buttonMultiplier = rpio.Pin(16)

func buttonMultiplierFunc(cfg *config.SafeConfig) {
	switch currentMultiplier := cfg.Multiplier; currentMultiplier {
	case 0.1:
		cfg.Multiplier = 0.5
	case 0.5:
		cfg.Multiplier = 1
	case 1:
		cfg.Multiplier = 2
	case 2:
		cfg.Multiplier = 5
	default:
		cfg.Multiplier = 0.1
	}
	if cfg.SSD1306 {
		ssd1306.Refresh(cfg)
	}
	config.Save(&cfg.Config)
}
