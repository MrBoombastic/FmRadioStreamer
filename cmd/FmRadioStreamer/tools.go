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
	"time"
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
	leds.Clear()
	tools.StopGPIO()
	time.Sleep(300 * time.Millisecond)
	core.SuperKill()
	time.Sleep(300 * time.Millisecond)
	oled.StopScreen()
	//time.Sleep(300 * time.Millisecond)
	//time.Sleep(300 * time.Millisecond)
	/*err := tools.StopGPIO()
	if err != nil {
		fmt.Println(err)
	}*/
}
