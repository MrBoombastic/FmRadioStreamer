package buttons

import (
	"context"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/stianeikeland/go-rpio/v4"
	"sync"
	"time"
)

// Init sets up physical buttons.
func Init() {
	buttons := [4]rpio.Pin{buttonDown, buttonUp, buttonMultiplier, buttonInvert}
	for _, item := range buttons {
		item.Input()
		item.PullUp()
		item.Detect(rpio.RiseEdge)
	}
}

var buttonFuncs = []func(cfg *config.SafeConfig){buttonDownFunc, buttonUpFunc, buttonMultiplierFunc, buttonInvertFunc}

// Listen enables frequency, multipliier and screen invertion controlling with physical buttons.
func Listen(wg *sync.WaitGroup, ctx context.Context, cfg *config.SafeConfig) {
	buttons := [4]rpio.Pin{buttonDown, buttonUp, buttonMultiplier, buttonInvert}
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Millisecond * 1000):
			for i, item := range buttons {
				if item.EdgeDetected() {
					cfg.Lock()
					fmt.Println(i)
					buttonFuncs[i](cfg)
					cfg.Unlock()
				}
			}
		}
	}
}
