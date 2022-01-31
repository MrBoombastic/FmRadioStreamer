package main

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/buttons"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	oled "github.com/MrBoombastic/FmRadioStreamer/pkg/screen"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"os"
	"periph.io/x/periph/devices/ssd1306"
	"time"
)

//var cfg = config.Get()
var multiplier = 0.1
var screen *ssd1306.Dev

func main() {
	// Init and start leds
	tools.InitGPIO()
	leds.InitLeds()
	buttons.InitButtons()
	go leds.QuadGreensLoopStart()
	//go leds.BlueLedLoopStart()
	// Init screen
	var err error
	screen, err = oled.Create()
	if err != nil {
		fmt.Println(err)
	}
	// Listen for process killing/exiting
	go StopApplicationHandler(screen)

	oled.RefreshScreen(screen)
	go buttons.ListenButtons(screen)

	// Code here!

	time.Sleep(90 * time.Second)
	StopPeriphs(screen)
	os.Exit(0)
}
