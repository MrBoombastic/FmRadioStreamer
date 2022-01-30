package screen

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"log"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/devices/ssd1306"
	"periph.io/x/periph/devices/ssd1306/image1bit"
	"periph.io/x/periph/host"
)

var cfg = config.Get()
var Frequency float64
var MiniMessage string
var Multiplier float64

func writer(img *image1bit.VerticalLSB, x int, y int, s string) {
	drawer := font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{C: image1bit.On},
		Face: basicfont.Face7x13,
		Dot:  fixed.P(x, y)}

	drawer.DrawString(s)
}

func initScreen() (*ssd1306.Dev, error) {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	b, err := i2creg.Open("1")
	if err != nil {
		log.Fatal(err)
	}
	//defer b.Close()

	return ssd1306.NewI2C(b, &ssd1306.DefaultOpts)

}

func CreateScreen() (*ssd1306.Dev, *image1bit.VerticalLSB) {
	screen, err := initScreen()
	if err != nil {
		fmt.Println(err)
	}
	return screen, image1bit.NewVerticalLSB(screen.Bounds())
}

func Draw(screen *ssd1306.Dev, img *image1bit.VerticalLSB) {
	if err := screen.Draw(screen.Bounds(), img, image.Point{}); err != nil {
		log.Fatal(err)
	}
}

func clearScreen(screen *ssd1306.Dev) {
	screen.StopScroll()
	if err := screen.Draw(screen.Bounds(), image1bit.NewVerticalLSB(screen.Bounds()), image.Point{}); err != nil {
		log.Fatal(err)
	}
}

func FillScreen(screen *ssd1306.Dev, img *image1bit.VerticalLSB) {
	clearScreen(screen)
	writer(img, 2, 11, cfg.PS)
	writer(img, 99, 11, fmt.Sprintf("%.1fx", Multiplier))
	maxRT := 15
	if len(cfg.RT) < 16 {
		maxRT = len(cfg.RT) - 1
	}
	writer(img, 1, 32, cfg.RT[0:maxRT])
	maxRT = 31
	if len(cfg.RT) < 32 {
		maxRT = len(cfg.RT) - 1
	}
	writer(img, 1, 42, cfg.RT[16:maxRT])
	writer(img, 71, 62, fmt.Sprintf("%.1f FM", Frequency))
	writer(img, 2, 62, MiniMessage)
	Draw(screen, img)

	err := screen.Scroll(ssd1306.Left, ssd1306.FrameRate5, 16, 48)
	if err != nil {
		return
	}
	//screen.Invert(true)
	//screen.Halt()
}
