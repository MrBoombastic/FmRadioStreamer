package leds

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

// Blink turns on selected LED and after some time turns it off.
func Blink(led rpio.Pin, duration time.Duration) {
	led.High()
	time.Sleep(duration)
	led.Low()
}

// Init sets up all needed LEDs.
func Init() {
	greenLed1.Output()
	greenLed2.Output()
	greenLed3.Output()
	greenLed4.Output()
	blueLed.Output()
	yellowLed.Output()
}
