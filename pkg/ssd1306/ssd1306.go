package ssd1306

import (
	"context"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/ssd1306"
	"periph.io/x/devices/v3/ssd1306/image1bit"
	"periph.io/x/host/v3"
	"strings"
	"sync"
	"time"
)

var screenConnection i2c.BusCloser
var img *image1bit.VerticalLSB

// screen is an open handle to the display controller.
var screen *ssd1306.Dev

// inverted defines wheter screen colours are normal or reverted (blue-on-black or black-on-blue).
var inverted = false

// Invert inverts screen colours.
func Invert() error {
	inverted = !inverted
	err := screen.Invert(inverted)
	if err != nil {
		return err
	}
	return nil
}

// writer writes text on img.
func writer(x int, y int, s string) {
	drawer := font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{C: image1bit.On},
		Face: basicfont.Face7x13,
		Dot:  fixed.P(x, y)}
	drawer.DrawString(s)
}

// Init sets up screen and handles shutting it down.
func Init(wg *sync.WaitGroup, ctx context.Context, cfg *config.SafeConfig) error {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			err := screen.StopScroll()
			if err != nil {
				return err
			}
			createImg()
			draw()
			err = screen.Invert(false)
			if err != nil {
				return err
			}
			err = screenConnection.Close()
			if err != nil {
				return err
			}
			return nil

		case <-time.After(time.Second):
			if screen == nil {
				_, err := host.Init()
				if err != nil {
					return err
				}
				screenConnection, err = i2creg.Open("1")
				if err != nil {
					return err
				}
				scr, err := ssd1306.NewI2C(screenConnection, &ssd1306.DefaultOpts)
				screen = scr
				Refresh(cfg)
			}
		}
	}
}

// createImg creates empty img.
func createImg() {
	img = image1bit.NewVerticalLSB(screen.Bounds())
}

// draw draws img on screen.
func draw() {
	if img == nil {
		createImg()
	}
	if err := screen.Draw(screen.Bounds(), img, image.Point{}); err != nil {
		logs.FmRadStrFatal(err)
	}
}

// MiniMessage shows custom message in bottom-left corer of screen for 2 seconds.
func MiniMessage(message string, cfg *config.SafeConfig) {
	if screen == nil {
		return
	}
	_ = screen.StopScroll()
	for x := 2; x <= 90; x++ {
		for y := 49; y <= 63; y++ {
			img.Set(x, y, image1bit.Off)
		}
	}
	writer(2, 62, message)
	draw()
	time.Sleep(2 * time.Second)
	Refresh(cfg)
}

// Refresh draws every possible element on the screen.
func Refresh(cfg *config.SafeConfig) {
	if screen == nil {
		return
	}
	_ = screen.StopScroll()
	createImg()
	writer(2, 11, cfg.PS)
	writer(71, 11, fmt.Sprintf("%.1f FM", cfg.Frequency))
	maxRT := 15
	if len(cfg.RT) < 16 {
		maxRT = len(cfg.RT)
	}
	writer(0, 32, cfg.RT[0:maxRT])
	maxRT = 31
	if len(cfg.RT) < 32 {
		maxRT = len(cfg.RT)
	}
	writer(0, 42, cfg.RT[16:maxRT])
	writer(99, 62, fmt.Sprintf("%.1fx", cfg.Multiplier))

	ip := strings.Split(tools.GetLocalIP(), ".")[2:4]
	writer(0, 62, fmt.Sprintf(".%v:%d", strings.Join(ip, "."), cfg.Port))

	draw()

	err := screen.Scroll(ssd1306.Left, ssd1306.FrameRate25, 16, 48)
	if err != nil {
		logs.FmRadStrError(err)
	}
}
