package ssd1306

import (
	"context"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"log"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/devices/ssd1306"
	"periph.io/x/periph/devices/ssd1306/image1bit"
	"periph.io/x/periph/host"
	"strings"
	"sync"
	"time"
)

var screenConnection i2c.BusCloser
var img *image1bit.VerticalLSB

// Screen is an open handle to the display controller
var Screen *ssd1306.Dev

// Inverted defines wheter screen colours are normal or reverted (blue-on-black or black-on-blue)
var Inverted = false

// writer writes text on img
func writer(x int, y int, s string) {
	drawer := font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{C: image1bit.On},
		Face: basicfont.Face7x13,
		Dot:  fixed.P(x, y)}

	drawer.DrawString(s)
}

// Init sets up screen and handles shutting it down
func Init(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			Screen.StopScroll()
			createImg()
			draw()
			Screen.Invert(false)
			screenConnection.Close()
			return nil

		case <-time.After(time.Second):
			if Screen == nil {
				_, err := host.Init()
				if err != nil {
					return err
				}
				screenConnection, err = i2creg.Open("1")
				if err != nil {
					return err
				}
				scr, err := ssd1306.NewI2C(screenConnection, &ssd1306.DefaultOpts)
				Screen = scr
				Refresh()
			}
		}
	}
}

// createImg creates empty img
func createImg() {
	img = image1bit.NewVerticalLSB(Screen.Bounds())
}

// draw draws img on Screen
func draw() {
	if img == nil {
		createImg()
	}
	if err := Screen.Draw(Screen.Bounds(), img, image.Point{}); err != nil {
		log.Fatal(err)
	}
}

// MiniMessage shows custom message in bottom-left corer of screen for 2 seconds
func MiniMessage(message string) {
	if Screen == nil {
		return
	}
	Screen.StopScroll()
	for x := 2; x <= 90; x++ {
		for y := 49; y <= 63; y++ {
			img.Set(x, y, image1bit.Off)
		}
	}
	writer(2, 62, message)
	draw()
	time.Sleep(2 * time.Second)
	Refresh()
}

// Refresh draws every possible element on the screen
func Refresh() {
	if Screen == nil {
		return
	}
	cfg := config.Get()
	Screen.StopScroll()
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

	ip := strings.Split(tools.LocalIP.String(), ".")[2:4]
	writer(0, 62, fmt.Sprintf(".%v:%d", strings.Join(ip, "."), cfg.Port))

	draw()

	err := Screen.Scroll(ssd1306.Left, ssd1306.FrameRate25, 16, 48)
	if err != nil {
		log.Println(err)
	}
}
