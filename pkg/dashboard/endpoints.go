package dashboard

import (
	"github.com/MrBoombastic/FmRadioStreamer/pkg/condlers"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/core"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/ssd1306"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"github.com/gofiber/fiber/v2"
	"log"
)

// music endpoint returns contents of `music` directory
func music(ctx *fiber.Ctx) {
	filesSlice, err := condlers.MusicDir()
	if err != nil {
		err := ctx.SendStatus(500)
		if err != nil {
			log.Println(err)
		}
	}
	err = ctx.JSON(filesSlice)
	if err != nil {
		log.Println(err)
	}
}

// stop endpoint plays silence
func stop(ctx *fiber.Ctx) {
	core.Play("")
	err := ctx.SendStatus(200)
	if err != nil {
		log.Println(err)
	}
}

// offair kill PiFmAdv entirely
func offair(ctx *fiber.Ctx) {
	core.Kill()
	err := ctx.SendStatus(200)
	if err != nil {
		log.Println(err)
	}
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
		log.Println(err)
		err := ctx.SendStatus(500)
		if err != nil {
			log.Println(err)
		}
		return
	}
	if search == "true" {
		err := ctx.JSON(result.Items[0])
		if err != nil {
			log.Println(err)
		}
	} else {
		err := ctx.SendStatus(200)
		if err != nil {
			log.Println(err)
		}
		leds.BlueLedEnabled = true
		err = condlers.DownloadWav("https://youtu.be/" + result.Items[0].ID.VideoID)
		leds.BlueLedEnabled = false
		if err != nil {
			log.Println(err)
		}
	}
}

// download endpoint performs downloading audio from other sites
func download(ctx *fiber.Ctx) {
	query := ctx.Query("q")
	err := ctx.SendStatus(200)
	if err != nil {
		log.Println(err)
	}
	leds.BlueLedEnabled = true
	err = condlers.DownloadWav(query)
	leds.BlueLedEnabled = false
	if err != nil {
		log.Println(err)
	}
}

// play endpoint plays selected file
func play(ctx *fiber.Ctx) {
	query := ctx.Query("q")
	core.Play(query)
	err := ctx.SendStatus(200)
	if err != nil {
		log.Println(err)
	}
}

// save endpoint updates current config
func save(ctx *fiber.Ctx) {
	newConfig := new(config.Config)
	if err := ctx.BodyParser(newConfig); err != nil {
		log.Fatalln("ERROR: Failed to parse new config!")
	}
	config.Save(*newConfig)
	if config.GetSSD1306() {
		ssd1306.Refresh()
	}
	err := ctx.SendStatus(200)
	if err != nil {
		log.Println(err)
	}
}

// configuration endpoint returns current config
func configuration(ctx *fiber.Ctx) {
	configMap := config.GetMap()
	err := ctx.JSON(configMap)
	if err != nil {
		log.Println(err)
	}
}
