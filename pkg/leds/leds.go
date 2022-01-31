package leds

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

func Blink(led rpio.Pin, miliseconds time.Duration) {
	for x := 0; x < 2; x++ {
		led.Toggle()
		time.Sleep(miliseconds)
	}
}
func InitLeds() {
	greenLed1.Output()
	greenLed2.Output()
	greenLed3.Output()
	greenLed4.Output()
	blueLed.Output()
	yellowLed.Output()
	ClearLeds()
}

func ClearLeds() {
	QuadGreensLoopStop()
	BlueLedLoopStop()
	yellowLed.Low()
}
