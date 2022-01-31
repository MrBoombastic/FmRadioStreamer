package buttons

import (
	"fmt"
	_ "fmt"
	"github.com/stianeikeland/go-rpio/v4"
	_ "github.com/stianeikeland/go-rpio/v4"
	"time"
	_ "time"
)

var (
	buttonUp         = rpio.Pin(20)
	buttonDown       = rpio.Pin(21)
	buttonSet        = rpio.Pin(16)
	buttonMultiplier = rpio.Pin(12)
	buttonTest       = rpio.Pin(14)
)

func InitButtons() {
	buttonUp.Input()
	buttonDown.Input()
	buttonSet.Input()
	buttonMultiplier.Input()
}

func ListenButtons() {
	buttonTest.Input()
	buttonTest.PullUp()
	buttonTest.Detect(rpio.FallEdge) // enable falling edge event detection

	fmt.Println("press a button")

	for i := 0; i < 2; {
		if buttonTest.EdgeDetected() { // check if event occured
			fmt.Println("button pressed")
			i++
		}
		time.Sleep(time.Second / 2)
	}
	buttonTest.Detect(rpio.NoEdge) // disable edge event detection
}
