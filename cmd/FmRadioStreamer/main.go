package main

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/buttons"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/core"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/dashboard"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	oled "github.com/MrBoombastic/FmRadioStreamer/pkg/screen"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"os"
	"periph.io/x/periph/devices/ssd1306"
	"time"
)

var screen *ssd1306.Dev

func main() {

	if os.Geteuid() != 0 {
		fmt.Println("Not running as sudo! Preventing system crash! Exiting...")
		os.Exit(0)
		return
	}
	// Init and start leds
	fmt.Println("Preparing peripherals...")
	tools.InitGPIO()
	leds.Init()
	buttons.Init()
	fmt.Println("Done preparing peripherals!")

	time.Sleep(time.Second * 2)

	fmt.Println("Starting peripherals...")
	go leds.QuadGreensLoopStart()
	// Init screen
	var err error
	err = oled.Create()
	if err != nil {
		fmt.Println(err)
	}

	core.Play("")

	// Listen for process killing/exiting
	go StopApplicationHandler()
	oled.RefreshScreen()
	go buttons.Listen()
	fmt.Println("Done starting peripherals!")
	go dashboard.Init()
	time.Sleep(1 * time.Hour)
	StopPeriphs()
	os.Exit(0)
}
