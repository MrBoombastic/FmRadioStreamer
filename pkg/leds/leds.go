package leds

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

func Blink(led rpio.Pin, duration time.Duration) {
	led.High()
	time.Sleep(duration)
	led.Low()
}

func Init() {
	greenLed1.Output()
	greenLed2.Output()
	greenLed3.Output()
	greenLed4.Output()
	blueLed.Output()
	yellowLed.Output()
}
