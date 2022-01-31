package buttons

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	oled "github.com/MrBoombastic/FmRadioStreamer/pkg/screen"
	"github.com/stianeikeland/go-rpio/v4"
	"math"
	"periph.io/x/periph/devices/ssd1306"
	"time"
)

var (
	buttonUp         = rpio.Pin(20)
	buttonDown       = rpio.Pin(21)
	buttonSet        = rpio.Pin(16)
	buttonMultiplier = rpio.Pin(12)
)

func Init() {
	buttons := [4]rpio.Pin{buttonDown, buttonUp, buttonSet, buttonMultiplier}
	for _, item := range buttons {
		item.Input()
		item.PullUp()
		item.Detect(rpio.AnyEdge)
	}
}

func Listen(screen *ssd1306.Dev) {
	for true {
		buttons := [4]rpio.Pin{buttonDown, buttonUp, buttonSet, buttonMultiplier}
		for i, item := range buttons {
			if item.EdgeDetected() {
				if i == 0 {
					currentFrequency := config.GetFrequency()
					currentMultiplier := config.GetMultiplier()
					if currentFrequency-currentMultiplier <= 87.2 {
						oled.MiniMessage = "MIN"
						leds.YellowBlink()
					} else {
						config.UpdateFrequency(math.Floor((config.GetFrequency()-config.GetMultiplier())*10) / 10)
					}
				}
				if i == 1 {
					currentFrequency := config.GetFrequency()
					currentMultiplier := config.GetMultiplier()
					if currentFrequency+currentMultiplier >= 108.9 {
						oled.MiniMessage = "MAX"
						go leds.YellowBlink()
					} else {
						config.UpdateFrequency(math.Floor((config.GetFrequency()+config.GetMultiplier())*10) / 10)
					}
				}
				if i == 3 {
					switch currentMultiplier := config.GetMultiplier(); currentMultiplier {
					case 0.1:
						config.UpdateMultiplier(0.5)
					case 0.5:
						config.UpdateMultiplier(1)
					case 1:
						config.UpdateMultiplier(2)
					case 2:
						config.UpdateMultiplier(5)
					default:
						config.UpdateMultiplier(0.1)
					}
				}
				fmt.Println(i)
				oled.RefreshScreen(screen)
			}
		}
		time.Sleep(time.Millisecond * 1000)
	}
}
