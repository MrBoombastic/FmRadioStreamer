package buttons

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

var (
	buttonUp         = rpio.Pin(20)
	buttonDown       = rpio.Pin(21)
	buttonSet        = rpio.Pin(16)
	buttonMultiplier = rpio.Pin(12)
)

func InitButtons() {
	buttons := [4]rpio.Pin{buttonUp, buttonDown, buttonSet, buttonMultiplier}
	for _, item := range buttons {
		item.PullUp()
		item.Detect(rpio.FallEdge)
	}
}

func ListenButtons() {
	for true {
		if buttonDown.EdgeDetected() { // check if event occured
			fmt.Println("down")
		}
		if buttonUp.EdgeDetected() { // check if event occured
			fmt.Println("up")
		}
		if buttonSet.EdgeDetected() { // check if event occured
			fmt.Println("set")
		}
		if buttonMultiplier.EdgeDetected() { // check if event occured
			fmt.Println("mult")
		}
		time.Sleep(time.Millisecond * 1000)
	}
}
