package leds

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

func Blink(led rpio.Pin, duration time.Duration) {
	for x := 0; x < 2; x++ {
		led.Toggle()
		time.Sleep(duration)
	}
}
func Init() {
	greenLed1.Output()
	greenLed2.Output()
	greenLed3.Output()
	greenLed4.Output()
	blueLed.Output()
	yellowLed.Output()
	Clear()
}

func Clear() {
	QuadGreensLoopStop()
	BlueLedLoopStop()
	yellowLed.Low()
}
