package buttons

import (
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/ssd1306"
	"github.com/stianeikeland/go-rpio/v4"
	"math"
)

var buttonDown = rpio.Pin(21)

func buttonDownFunc(cfg *config.SafeConfig) {
	if cfg.Frequency-cfg.Multiplier < 76.0 {
		go leds.YellowBlink()
		if cfg.SSD1306 {
			ssd1306.MiniMessage("MIN FREQ!", cfg)
		}
	} else {
		cfg.Frequency = math.Floor((cfg.Frequency-cfg.Multiplier)*10) / 10
		config.Save(&cfg.Config)
		if cfg.SSD1306 {
			ssd1306.Refresh(cfg)
		}
	}
}
