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
	screen, img := oled.CreateScreen()
	oled.Multiplier = multiplier
	oled.Frequency = cfg.Frequency
	oled.FillScreen(screen, img)
	time.Sleep(3 * time.Second)
	oled.MiniMessage = "XDDDDD"
	oled.FillScreen(screen, img)
	time.Sleep(3 * time.Second)
	oled.MiniMessage = "OK"
	oled.FillScreen(screen, img)
	time.Sleep(1 * time.Hour)
}
