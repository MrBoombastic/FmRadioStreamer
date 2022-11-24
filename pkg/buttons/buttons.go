package buttons

import (
	"context"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/ssd1306"
	"github.com/stianeikeland/go-rpio/v4"
	"math"
	"sync"
	"time"
)

var (
	buttonUp         = rpio.Pin(20)
	buttonDown       = rpio.Pin(21)
	buttonMultiplier = rpio.Pin(16)
	buttonInvert     = rpio.Pin(12)
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
					if i == 0 {
						cfg.Lock()
						if cfg.Frequency-cfg.Multiplier < 76.0 {
							go leds.YellowBlink()
							if cfg.SSD1306 {
								ssd1306.MiniMessage("MIN FREQ!", cfg)
							}
						} else {
							cfg.Frequency = math.Floor((cfg.Frequency-cfg.Multiplier)*10) / 10
							config.Save(&cfg.Config)
							if cfg.SSD1306 {
								ssd1306.Refresh(cfg)
							}
						}
						cfg.Unlock()
					}
					if i == 1 {
						cfg.Lock()
						if cfg.Frequency+cfg.Multiplier > 108.0 {
							go leds.YellowBlink()
							if cfg.SSD1306 {
								ssd1306.MiniMessage("MAX FREQ!", cfg)
							}
						} else {
							cfg.Frequency = math.Floor((cfg.Frequency-cfg.Multiplier)*10) / 10
							config.Save(&cfg.Config)
							if cfg.SSD1306 {
								ssd1306.Refresh(cfg)
							}
						}
						cfg.Unlock()
					}
					if i == 2 {
						cfg.Lock()
						switch currentMultiplier := cfg.Multiplier; currentMultiplier {
						case 0.1:
							cfg.Multiplier = 0.5
						case 0.5:
							cfg.Multiplier = 1
						case 1:
							cfg.Multiplier = 2
						case 2:
							cfg.Multiplier = 5
						default:
							cfg.Multiplier = 0.1
						}
						if cfg.SSD1306 {
							ssd1306.Refresh(cfg)
						}
						config.Save(&cfg.Config)
						cfg.Unlock()
					}
					if i == 3 {
						cfg.Lock()
						if cfg.SSD1306 {
							err := ssd1306.Invert()
							if err != nil {
								logs.FmRadStrError(err)
							}
						}
						cfg.Unlock()
					}
				}
			}
		}
	}
}
