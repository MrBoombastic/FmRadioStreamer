package tools

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"os"
	"os/signal"
	"syscall"
)

func EndHandler() {
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
	//	oled.FillScreen(0.5)
	//err := screen.StopScreen()
	//if err != nil {
	//	fmt.Println(err)
	//}
	os.Exit(1)
}
