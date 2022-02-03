package main

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/core"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	oled "github.com/MrBoombastic/FmRadioStreamer/pkg/screen"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"os"
	"os/signal"
	"periph.io/x/periph/devices/ssd1306"
	"syscall"
)

func StopApplicationHandler(screen *ssd1306.Dev) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		<-sigs
		done <- true
	}()
	<-done
	fmt.Println()
	fmt.Println("Exiting...")
	oled.StopScreen()
	StopPeriphs()
	os.Exit(0)
}

func StopPeriphs() {
	leds.Clear()
	oled.StopScreen()
	tools.StopGPIO()
	core.Kill()
}
