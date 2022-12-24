package buttons

import (
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/ssd1306"
	"github.com/stianeikeland/go-rpio/v4"
	"math"
)

var buttonUp = rpio.Pin(20)

func buttonUpFunc(cfg *config.SafeConfig) {
	if cfg.Frequency+cfg.Multiplier > 108.0 {
		go leds.YellowBlink()
		if cfg.SSD1306 {
			ssd1306.MiniMessage("MAX FREQ!", cfg)
		}
	} else {
		cfg.Frequency = math.Floor((cfg.Frequency+cfg.Multiplier)*10) / 10
		config.Save(&cfg.Config)
		if cfg.SSD1306 {
			ssd1306.Refresh(cfg)
		}
	}
}
