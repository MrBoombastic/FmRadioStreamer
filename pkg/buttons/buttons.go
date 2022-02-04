package buttons

import (
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	oled "github.com/MrBoombastic/FmRadioStreamer/pkg/screen"
	"github.com/stianeikeland/go-rpio/v4"
	"math"
	"time"
)

var (
	buttonUp         = rpio.Pin(20)
	buttonDown       = rpio.Pin(21)
	buttonMultiplier = rpio.Pin(16)
	buttonInvert     = rpio.Pin(12)
)

func Init() {
	buttons := [4]rpio.Pin{buttonDown, buttonUp, buttonMultiplier, buttonInvert}
	for _, item := range buttons {
		item.Input()
		item.PullUp()
		item.Detect(rpio.AnyEdge)
	}
}

func Listen() {
	for true {
		buttons := [4]rpio.Pin{buttonDown, buttonUp, buttonMultiplier, buttonInvert}
		for i, item := range buttons {
			if item.EdgeDetected() {
				if i == 0 {
					currentFrequency := config.GetFrequency()
					currentMultiplier := config.GetMultiplier()
					if currentFrequency-currentMultiplier < 76.0 {
						oled.MiniMessage = "MIN"
						leds.YellowBlink()
					} else {
						config.UpdateFrequency(math.Floor((config.GetFrequency()-config.GetMultiplier())*10) / 10)
					}
				}
				if i == 1 {
					currentFrequency := config.GetFrequency()
					currentMultiplier := config.GetMultiplier()
					if currentFrequency+currentMultiplier > 108.0 {
						oled.MiniMessage = "MAX"
						go leds.YellowBlink()
					} else {
						config.UpdateFrequency(math.Floor((config.GetFrequency()+config.GetMultiplier())*10) / 10)
					}
				}
				if i == 2 {
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
				if i == 3 {
					oled.ScreenInverted = !oled.ScreenInverted
					oled.Screen.Invert(oled.ScreenInverted)
				}
				// Do not refresh screen, when only inverting colors
				if i < 3 {
					oled.RefreshScreen()
				}
			}
		}
		time.Sleep(time.Millisecond * 1000)
	}
}