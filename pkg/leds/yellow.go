package leds

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

var yellowLed = rpio.Pin(5)

// YellowBlink blinks with yellow LED (Blink wrapper).
func YellowBlink() {
	Blink(yellowLed, 2*time.Second)
}
