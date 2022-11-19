package dashboard

import (
	"github.com/MrBoombastic/FmRadioStreamer/pkg/condlers"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/core"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/ssd1306"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/pbar1/pkill-go"
	"log"
	"os"
)

// music endpoint returns contents of `music` directory
func music(ctx *fiber.Ctx) {
	filesSlice, err := condlers.MusicDir()
	if err != nil {
		_ = ctx.SendStatus(500)
	}
	err = ctx.JSON(filesSlice)
	if err != nil {
		logs.FmRadStrError(err)
	}
}

// stop endpoint plays silence
func stop(ctx *fiber.Ctx) {
	core.Play("")
	_ = ctx.SendStatus(200)
}

// offair kill PiFmAdv entirely
func offair(ctx *fiber.Ctx) {
	_, err := pkill.Pkill("pi_fm_adv", os.Interrupt)
	if err != nil {
		logs.FmRadStrError(err)
		_ = ctx.SendStatus(500)
	}
	_ = ctx.SendStatus(200)
	if config.GetSSD1306() {
		ssd1306.MiniMessage("OFF-AIR")
	}
}

// yt endpoint performs searching or downloading audio from YouTube
func yt(ctx *fiber.Ctx) {
	search := ctx.Query("search")
	query := ctx.Query("q")
	result, err := tools.SearchYouTube(query)
	if err != nil {
		logs.FmRadStrError(err)
		_ = ctx.SendStatus(500)
		return
	}
	if search == "true" {
		err := ctx.JSON(result.Items[0])
		if err != nil {
			logs.FmRadStrError(err)
		}
	} else {
		_ = ctx.SendStatus(200)
		leds.BlueLedEnabled = true
		cfg := config.Get()
		err = condlers.Download("https://youtu.be/"+result.Items[0].ID.VideoID, cfg.Format)
		leds.BlueLedEnabled = false
		if err != nil {
			logs.FmRadStrError(err)
		}
	}
}

// youtubeDl endpoint performs downloading audio from other sites
func youtubeDl(ctx *fiber.Ctx) {
	query := ctx.Query("q")
	_ = ctx.SendStatus(200)
	leds.BlueLedEnabled = true
	cfg := config.Get()
	err := condlers.Download(query, cfg.Format)
	leds.BlueLedEnabled = false
	if err != nil {
		logs.FmRadStrError(err)
	}
}

// playFile endpoint plays selected file
func playFile(ctx *fiber.Ctx) {
	query := ctx.Query("q")
	err := core.Play(query)
	if err != nil {
		logs.PiFmAdvError(err)
	}
	_ = ctx.SendStatus(200)
}

// playStream endpoint plays remote file via SoX
func playStream(ctx *fiber.Ctx) {
	query := ctx.Query("q")
	err := core.Sox(query)
	if err != nil {
		logs.PiFmAdvError(err)
		_ = ctx.SendStatus(500)
	}
	_ = ctx.SendStatus(200)
}

// save endpoint updates current config
func save(ctx *fiber.Ctx) {
	newConfig := new(config.Config)
	if err := ctx.BodyParser(newConfig); err != nil {
		logs.PiFmAdvError(err)
		_ = ctx.SendStatus(500)

	}
	config.Save(*newConfig)
	if config.GetSSD1306() {
		ssd1306.Refresh()
	}
	_ = ctx.SendStatus(200)
}

// configuration endpoint returns current config
func configuration(ctx *fiber.Ctx) {
	configMap := config.GetMap()
	err := ctx.JSON(configMap)
	if err != nil {
		log.Println(err)
	}
}
