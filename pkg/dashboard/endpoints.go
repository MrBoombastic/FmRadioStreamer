package dashboard

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/condlers"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/core"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/ssd1306"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"github.com/pbar1/pkill-go"
	"os"
)

// music endpoint returns contents of `music` directory
func music(ctx *RadioContext) {
	filesSlice, err := condlers.MusicDir()
	if err != nil {
		_ = ctx.Fiber.SendStatus(500)
	}
	err = ctx.Fiber.JSON(filesSlice)
	if err != nil {
		logs.FmRadStrError(err)
	}
}

// stop endpoint plays silence
func stop(ctx *RadioContext) {
	err := core.Play(tools.Params{Type: tools.SilenceType})
	if err != nil {
		logs.PiFmAdvError(err)
		_ = ctx.Fiber.SendStatus(500)
	}
	_ = ctx.Fiber.SendStatus(200)
}

// offair kill PiFmAdv entirely
func offair(ctx *RadioContext) {
	_, err := pkill.Pkill("pi_fm_adv", os.Interrupt)
	if err != nil {
		logs.FmRadStrError(err)
		_ = ctx.Fiber.SendStatus(500)
	}
	_ = ctx.Fiber.SendStatus(200)
	ctx.Cfg.Lock()
	if ctx.Cfg.SSD1306 {
		ssd1306.MiniMessage("OFF-AIR")
	}
	ctx.Cfg.Unlock()
}

// yt endpoint performs searching or downloading audio from YouTube
func yt(ctx *RadioContext) {
	search := ctx.Fiber.Query("search")
	query := ctx.Fiber.Query("q")
	ctx.Cfg.Lock()
	key := ctx.Cfg.YouTubeAPIKey
	ctx.Cfg.Unlock()
	result, err := tools.SearchYouTube(query, key)
	if err != nil {
		logs.FmRadStrError(err)
		_ = ctx.Fiber.SendStatus(500)
		return
	}
	if search == "true" {
		err := ctx.Fiber.JSON(result.Items[0])
		if err != nil {
			logs.FmRadStrError(err)
		}
	} else {
		_ = ctx.Fiber.SendStatus(200)
		leds.BlueLedEnabled = true
		ctx.Cfg.Lock()
		err = condlers.Download("https://youtu.be/"+result.Items[0].ID.VideoID, ctx.Cfg.Format)
		ctx.Cfg.Unlock()
		leds.BlueLedEnabled = false
		if err != nil {
			logs.FmRadStrError(err)
		}
	}
}

// youtubeDl endpoint performs downloading audio from other sites
func youtubeDl(ctx *RadioContext) {
	query := ctx.Fiber.Query("q")
	_ = ctx.Fiber.SendStatus(200)
	leds.BlueLedEnabled = true
	ctx.Cfg.Lock()
	err := condlers.Download(query, ctx.Cfg.Format)
	ctx.Cfg.Unlock()
	leds.BlueLedEnabled = false
	if err != nil {
		logs.FmRadStrError(err)
	}
}

// playFile endpoint plays selected file
func playFile(ctx *RadioContext) {
	query := ctx.Fiber.Query("q")
	err := core.Play(tools.Params{Type: tools.FileType, Audio: query})
	if err != nil {
		logs.PiFmAdvError(err)
		_ = ctx.Fiber.SendStatus(500)
	}
	_ = ctx.Fiber.SendStatus(200)
}

// playStream endpoint plays remote file via SoX
func playStream(ctx *RadioContext) {
	query := ctx.Fiber.Query("q")
	err := core.Play(tools.Params{Type: tools.StreamType, Audio: query})
	if err != nil {
		logs.PiFmAdvError(err)
		_ = ctx.Fiber.SendStatus(500)
	}
	_ = ctx.Fiber.SendStatus(200)
}

// save endpoint updates current config
func save(ctx *RadioContext) {
	newCfg := new(config.SafeConfig)
	if err := ctx.Fiber.BodyParser(newCfg); err != nil {
		logs.PiFmAdvError(err)
		_ = ctx.Fiber.SendStatus(500)

	}
	ctx.Cfg.Lock()
	config.Save(newCfg)
	if ctx.Cfg.SSD1306 {
		ssd1306.Refresh()
	}
	ctx.Cfg.Unlock()
	_ = ctx.Fiber.SendStatus(200)
}

// configuration endpoint returns current config
func configuration(ctx *RadioContext) {
	ctx.Cfg.Lock()
	configMap := tools.ConfigToMap(ctx.Cfg)
	ctx.Cfg.Unlock()
	fmt.Println(configMap)
	err := ctx.Fiber.JSON(configMap)
	if err != nil {
		logs.FmRadStrError(err)
		_ = ctx.Fiber.SendStatus(500)
	}
}
