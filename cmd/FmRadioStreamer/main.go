package main

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	oled "github.com/MrBoombastic/FmRadioStreamer/pkg/screen"
	"periph.io/x/periph/devices/ssd1306"
	"time"
)

//var cfg = config.Get()
var multiplier = 0.1
var screen *ssd1306.Dev

func main() {
	// Init and start leds
	leds.InitLeds()
	go leds.QuadGreensLoopStart()
	go leds.BlueLedLoopStart()

	// Listen for process killing/exiting
	go endHandler()

	// Init screen
	var err error
	screen, err = oled.CreateScreen()
	if err != nil {
		fmt.Println(err)
	}
	oled.Multiplier = multiplier
	oled.FillScreen(screen)
	time.Sleep(3 * time.Second)
}
