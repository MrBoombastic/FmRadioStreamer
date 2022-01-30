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

var greenSleepInterval = 500 * time.Millisecond

func QuadGreensLoopStart() {
	greenLed4.High()
	for true {
		greenLed4.Toggle()
		greenLed1.Toggle()
		time.Sleep(greenSleepInterval)
		greenLed1.Toggle()
		greenLed2.Toggle()
		time.Sleep(greenSleepInterval)
		greenLed2.Toggle()
		greenLed3.Toggle()
		time.Sleep(greenSleepInterval)
		greenLed3.Toggle()
		greenLed4.Toggle()
		time.Sleep(greenSleepInterval)
	}
}
func QuadGreensLoopStop() {
	greenLed1.Low()
	greenLed2.Low()
	greenLed3.Low()
	greenLed4.Low()
}
