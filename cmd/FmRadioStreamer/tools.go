package main

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/core"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	oled "github.com/MrBoombastic/FmRadioStreamer/pkg/screen"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"os"
	"os/signal"
	"syscall"
)

func StopApplicationHandler() {
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
	StopPeriphs()
	os.Exit(0)
}

func StopPeriphs() {
	core.SuperKill()
	leds.Clear()
	oled.StopScreen()
	err := tools.StopGPIO()
	if err != nil {
		fmt.Println(err)
	}
}
