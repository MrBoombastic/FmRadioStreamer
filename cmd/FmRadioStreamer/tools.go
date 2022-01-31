package main

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	oled "github.com/MrBoombastic/FmRadioStreamer/pkg/screen"
	"os"
	"os/signal"
	"syscall"
)

func endHandler() {
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
	leds.ClearLeds()
	oled.StopScreen(screen)

	os.Exit(1)
}
