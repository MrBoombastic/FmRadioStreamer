package screen

import (
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
)

var cfg = config.Get()
var Frequency = cfg.Frequency
var MiniMessage string
var Multiplier float64
var screenConnection i2c.BusCloser

func writer(img *image1bit.VerticalLSB, x int, y int, s string) {
	drawer := font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{C: image1bit.On},
		Face: basicfont.Face7x13,
		Dot:  fixed.P(x, y)}

	drawer.DrawString(s)
}

func CreateScreen() (*ssd1306.Dev, error) {
	_, err := host.Init()
	if err != nil {
		return nil, err
	}

	screenConnection, err = i2creg.Open("1")
	if err != nil {
		log.Fatal(err)
	}
	return ssd1306.NewI2C(screenConnection, &ssd1306.DefaultOpts)
}

func createImg(screen *ssd1306.Dev) *image1bit.VerticalLSB {
	return image1bit.NewVerticalLSB(screen.Bounds())
}

func draw(screen *ssd1306.Dev, img *image1bit.VerticalLSB) {
	if err := screen.Draw(screen.Bounds(), img, image.Point{}); err != nil {
		log.Fatal(err)
	}
}

func FillScreen(screen *ssd1306.Dev) {
	screen.StopScroll()
	img := createImg(screen)
	writer(img, 2, 11, cfg.PS)
	writer(img, 71, 11, fmt.Sprintf("%.1f FM", Frequency))
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
	writer(img, 99, 62, fmt.Sprintf("%.1fx", Multiplier))
	if MiniMessage == "100" { //assuming this message is from FFmpeg
		MiniMessage = ""
	}
	if MiniMessage != "" {
		writer(img, 2, 62, MiniMessage)
	} else {
		ip := strings.Split(tools.GetLocalIP().String(), ".")[2:4]
		writer(img, 2, 62, fmt.Sprintf(".%v:%d", strings.Join(ip, "."), cfg.Port))
	}
	draw(screen, img)

	err := screen.Scroll(ssd1306.Left, ssd1306.FrameRate5, 16, 48)
	if err != nil {
		return
	}
	//screen.Invert(true)
}

func StopScreen(screen *ssd1306.Dev) {
	screen.StopScroll()
	draw(screen, createImg(screen))
	screenConnection.Close()
}
