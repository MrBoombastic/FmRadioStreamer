package leds

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

var blueLedEnabled = false
var blueLed = rpio.Pin(7)

func BlueLedLoopStart() {
	blueLedEnabled = true
	for true {
		blueLed.High()
		time.Sleep(500 * time.Millisecond)
		blueLed.Low()
		if blueLedEnabled == false {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func BlueLedLoopStop() {
	blueLedEnabled = false
	blueLed.Low()
}
