package main

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/buttons"
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
	screen, err = oled.Create()
	if err != nil {
		fmt.Println(err)
	}
	// Listen for process killing/exiting
	go StopApplicationHandler(screen)
	oled.RefreshScreen(screen)
	go buttons.Listen(screen)
	fmt.Println("Done starting peripherals!")
	go dashboard.Init()
	time.Sleep(90 * time.Second)
	StopPeriphs(screen)
	os.Exit(0)
}
