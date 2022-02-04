package screen

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

var MiniMessage string
var screenConnection i2c.BusCloser
var Screen *ssd1306.Dev
var ScreenInverted = false

func writer(img *image1bit.VerticalLSB, x int, y int, s string) {
	drawer := font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{C: image1bit.On},
		Face: basicfont.Face7x13,
		Dot:  fixed.P(x, y)}

	drawer.DrawString(s)
}

func Create(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			Screen.StopScroll()
			draw(createImg())
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
				RefreshScreen()
			}
		}
	}
}

func createImg() *image1bit.VerticalLSB {
	return image1bit.NewVerticalLSB(Screen.Bounds())
}

func draw(img *image1bit.VerticalLSB) {
	if err := Screen.Draw(Screen.Bounds(), img, image.Point{}); err != nil {
		log.Fatal(err)
	}
}

func RefreshScreen() {
	if Screen == nil {
		return
	}
	cfg := config.Get()
	Screen.StopScroll()
	img := createImg()
	writer(img, 2, 11, cfg.PS)
	writer(img, 71, 11, fmt.Sprintf("%.1f FM", cfg.Frequency))
	maxRT := 15
	if len(cfg.RT) < 16 {
		maxRT = len(cfg.RT)
	}
	writer(img, 1, 32, cfg.RT[0:maxRT])
	maxRT = 31
	if len(cfg.RT) < 32 {
		maxRT = len(cfg.RT)
	}
	writer(img, 1, 42, cfg.RT[16:maxRT])
	writer(img, 99, 62, fmt.Sprintf("%.1fx", cfg.Multiplier))
	if MiniMessage == "100" { //assuming this message is from FFmpeg
		MiniMessage = ""
	}
	if MiniMessage != "" {
		writer(img, 2, 62, MiniMessage)
	} else {
		ip := strings.Split(tools.GetLocalIP().String(), ".")[2:4]
		writer(img, 2, 62, fmt.Sprintf(".%v:%d", strings.Join(ip, "."), cfg.Port))
	}
	draw(img)

	err := Screen.Scroll(ssd1306.Left, ssd1306.FrameRate5, 16, 48)
	if err != nil {
		return
	}
	//screen.Invert(true)
}
