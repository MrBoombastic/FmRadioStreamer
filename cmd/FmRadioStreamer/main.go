package main

import (
	"context"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/buttons"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/core"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/dashboard"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	oled "github.com/MrBoombastic/FmRadioStreamer/pkg/screen"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"os"
	"os/signal"
	"periph.io/x/periph/devices/ssd1306"
	"sync"
	"syscall"
)

var screen *ssd1306.Dev

func main() {
	if os.Geteuid() != 0 {
		fmt.Println("Not running as sudo! Preventing system crash! Exiting...")
		os.Exit(0)
		return
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	var wg sync.WaitGroup

	fmt.Println("Starting peripherals...")
	tools.InitGPIO()

	leds.Init()
	wg.Add(1)
	go leds.QuadGreensLoopStart(&wg, ctx)
	wg.Add(1)
	go leds.BlueLedLoopStart(&wg, ctx)

	// Init screen
	wg.Add(1)
	go oled.Create(&wg, ctx)

	// Init buttons
	buttons.Init()
	wg.Add(1)
	go buttons.Listen(&wg, ctx)

	dashboard.Init()

	core.Play("")

	wg.Wait()
	fmt.Println("Exiting...")
	fmt.Println("gpio")
	//tools.StopGPIO()
	fmt.Println("core")
	core.SuperKill()
	fmt.Println("Exited")
}
