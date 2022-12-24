package buttons

import (
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/ssd1306"
	"github.com/stianeikeland/go-rpio/v4"
)

var buttonInvert = rpio.Pin(12)

func buttonInvertFunc(cfg *config.SafeConfig) {
	if cfg.SSD1306 {
		err := ssd1306.Invert()
		if err != nil {
			logs.FmRadStrError(err)
		}
	}
}
