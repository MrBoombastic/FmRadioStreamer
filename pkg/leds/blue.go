package leds

import (
	"context"
	"github.com/stianeikeland/go-rpio/v4"
	"sync"
	"time"
)

var BlueLedEnabled = false
var blueLedSleepInterval = 500 * time.Millisecond
var blueLed = rpio.Pin(7)

func BlueLedLoopStart(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			blueLed.Low()
			return
		case <-time.After(blueLedSleepInterval / 2):
			if !BlueLedEnabled {
				continue
			}
			blueLed.High()
			time.Sleep(blueLedSleepInterval)
			blueLed.Low()
			time.Sleep(blueLedSleepInterval / 2)
		}
	}
}
