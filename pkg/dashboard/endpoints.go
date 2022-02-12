package dashboard

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/condlers"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/core"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/ssd1306"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"github.com/gofiber/fiber/v2"
	"log"
)

func music(ctx *fiber.Ctx) {
	filesSlice, err := musicList()
	if err != nil {
		ctx.SendStatus(500)
	}
	ctx.JSON(filesSlice)
}

func loudstop(ctx *fiber.Ctx) {
	core.Play("")
	ctx.SendStatus(200)
}

func superstop(ctx *fiber.Ctx) {
	core.Kill()
	ctx.SendStatus(200)
}

func yt(ctx *fiber.Ctx) {
	search := ctx.Query("search")
	query := ctx.Query("q")
	result, err := tools.SearchYouTube(query)
	if err != nil {
		log.Println(err)
		ctx.SendStatus(500)
		return
	}
	if search == "true" {
		ctx.JSON(result.Items[0].Snippet)
	} else {
		ctx.SendStatus(200)
		leds.BlueLedEnabled = true
		err := condlers.DownloadAudioFromYoutube(result.Items[0].ID.VideoID, result.Items[0].Snippet.Title)
		leds.BlueLedEnabled = false
		if err != nil {
			log.Println(err)
		}
	}
}

func play(ctx *fiber.Ctx) {
	fmt.Println("play")
	query := ctx.Query("q")
	core.Play(query)
	ctx.SendStatus(200)
}

func save(ctx *fiber.Ctx) {
	newConfig := new(config.Config)
	if err := ctx.BodyParser(newConfig); err != nil {
		log.Fatalln("ERROR: Failed to parse new config!")
	}
	config.Save(*newConfig)
	if config.GetSSD1306() {
		ssd1306.Refresh()
	}
	ctx.SendStatus(200)
}

func configuration(ctx *fiber.Ctx) {
	configMap := config.GetMap()
	ctx.JSON(configMap)
}
