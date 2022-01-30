package main

import (
	"fmt"
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

	// Init screen
	screen, err := oled.CreateScreen()
	if err != nil {
		fmt.Println(err)
	}
	oled.Multiplier = multiplier
	oled.Frequency = cfg.Frequency
	oled.FillScreen(screen)
	time.Sleep(3 * time.Second)
	oled.Multiplier = 69
	oled.Frequency = 21.37
	oled.MiniMessage = "XDDDDD"
	oled.FillScreen(screen)
	time.Sleep(3 * time.Second)
	oled.Multiplier = 420
	oled.Frequency = 13.37
	oled.MiniMessage = "OK"
	oled.FillScreen(screen)
	time.Sleep(1 * time.Hour)
}
