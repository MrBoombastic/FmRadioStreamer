package leds

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

var blueLedEnabled = true
var blueLed = rpio.Pin(7)

func BlueLedLoopStart() {
	blueLedEnabled = true
	for true {
		if blueLedEnabled == false {
			break
		} else {
			Blink(blueLed, 500*time.Millisecond)
		}
	}
}

func BlueLedLoopStop() {
	blueLedEnabled = false
	blueLed.Low()
}
