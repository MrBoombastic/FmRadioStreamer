package main

import (
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	oled "github.com/MrBoombastic/FmRadioStreamer/pkg/screen"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"time"
)

var cfg = config.Get()
var multiplier = 0.1

func main() {
	// Reset leds state
	leds.InitLeds()
	go leds.QuadGreensLoopStart()
	go leds.BlueLedLoopStart()

	// Listen for process killing/exiting
	go tools.EndHandler()

	// Inits screen
	oled.FillScreen(cfg.Frequency, multiplier, "OK", false)
	time.Sleep(1 * time.Hour)
}
