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
	"strings"
)

var cfg = config.Get()

func writer(img *image1bit.VerticalLSB, x int, y int, s string) {
	drawer := font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{C: image1bit.On},
		Face: basicfont.Face7x13,
		Dot:  fixed.P(x, y)}

	drawer.DrawString(s)
}

func FillScreen(freq float32, multiplier float64, miniMessage string) {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	b, err := i2creg.Open("1")
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	screen, _ := ssd1306.NewI2C(b, &ssd1306.DefaultOpts)

	img := image1bit.NewVerticalLSB(screen.Bounds())
	writer(img, 2, 11, cfg.PS)
	writer(img, 99, 11, fmt.Sprintf("%.1fx", multiplier))

	splittedRT := strings.Split(cfg.RT, "")
	RTpart0 := ""
	RTpart1 := ""
	for i, char := range splittedRT {
		line := i / 16
		if line == 0 {
			RTpart0 += char
		} else if line == 1 {
			RTpart1 += char
		} else {
			break
		}
	}
	writer(img, 1, 32, RTpart0)
	writer(img, 1, 42, RTpart1)
	writer(img, 2, 62, miniMessage)
	writer(img, 71, 62, fmt.Sprintf("%.1f FM", freq))

	if err := screen.Draw(screen.Bounds(), img, image.Point{}); err != nil {
		log.Fatal(err)
	}

	err = screen.Scroll(ssd1306.Left, ssd1306.FrameRate2, 16, 48)
	if err != nil {
		return
	}
	//screen.Invert(true)
	//screen.Halt()
}
