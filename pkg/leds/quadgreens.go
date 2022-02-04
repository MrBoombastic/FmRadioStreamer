package leds

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

var (
	greenLed1 = rpio.Pin(19)
	greenLed2 = rpio.Pin(6)
	greenLed3 = rpio.Pin(13)
	greenLed4 = rpio.Pin(26)
)

var greensSleepInterval = 500 * time.Millisecond
var GreensLoopEnabled = true

func QuadGreensLoopStart() {
	GreensLoopEnabled = true
	greenLed4.High()
	for GreensLoopEnabled {
		greenLed4.Toggle()
		greenLed1.Toggle()
		time.Sleep(greensSleepInterval)
		greenLed1.Toggle()
		greenLed2.Toggle()
		time.Sleep(greensSleepInterval)
		greenLed2.Toggle()
		greenLed3.Toggle()
		time.Sleep(greensSleepInterval)
		greenLed3.Toggle()
		greenLed4.Toggle()
		time.Sleep(greensSleepInterval)
	}
}
