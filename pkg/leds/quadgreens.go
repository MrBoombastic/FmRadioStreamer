package leds

import (
	"context"
	"github.com/stianeikeland/go-rpio/v4"
	"sync"
	"time"
)

var (
	greenLed1 = rpio.Pin(19)
	greenLed2 = rpio.Pin(6)
	greenLed3 = rpio.Pin(13)
	greenLed4 = rpio.Pin(26)
)

var greensSleepInterval = 500 * time.Millisecond

// QuadGreensLoop handles quad green LEDs activity - just blinking in loop in circle
func QuadGreensLoop(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	greenLed4.High()
	for {
		select {
		case <-ctx.Done():
			greenLed1.Low()
			greenLed2.Low()
			greenLed3.Low()
			greenLed4.Low()
			return
		case <-time.After(greensSleepInterval / 2):
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
			time.Sleep(greensSleepInterval / 2)
		}
	}
}
