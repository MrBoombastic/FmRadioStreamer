package leds

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

var BlueLedEnabled = false
var blueLed = rpio.Pin(7)

func BlueLedLoopStart() {
	BlueLedEnabled = true
	for true {
		blueLed.High()
		time.Sleep(500 * time.Millisecond)
		blueLed.Low()
		if BlueLedEnabled == false {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}
