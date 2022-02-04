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
var greensLoopEnabled = true

func QuadGreensLoopStart() {
	greensLoopEnabled = true
	greenLed4.High()
	for true {
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
		if greensLoopEnabled == false {
			break
		}
	}
}
func QuadGreensLoopStop() {
	greensLoopEnabled = false
	greenLed1.Low()
	greenLed2.Low()
	greenLed3.Low()
	greenLed4.Low()
}
