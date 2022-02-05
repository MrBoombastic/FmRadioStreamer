package dashboard

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/condlers"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/core"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/screen"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"log"
)

var musicEndpoint = EndpointData{
	Endpoint: music,
}

func music(ctx Context) {
	filesSlice := musicList()
	ctx.Server.JSON(filesSlice)
}

var loudstopEndpoint = EndpointData{
	Endpoint: loudstop,
}

func loudstop(ctx Context) {
	core.Play("")
	ctx.Server.SendStatus(200)
}

var superstopEndpoint = EndpointData{
	Endpoint: superstop,
}

func superstop(ctx Context) {
	core.SuperKill()
	ctx.Server.SendStatus(200)
}

var ytEndpoint = EndpointData{
	Endpoint: yt,
}

func yt(ctx Context) {
	search := ctx.Server.Query("search")
	query := ctx.Server.Query("q")
	result := tools.SearchYouTube(query)
	if search == "true" {
		ctx.Server.JSON(result.Items[0].Snippet)
	} else {
		ctx.Server.SendStatus(200)
		leds.BlueLedEnabled = true
		err := condlers.DownloadAudioFromYoutube(result.Items[0].ID.VideoID, result.Items[0].Snippet.Title)
		leds.BlueLedEnabled = false
		if err != nil {
			fmt.Println(err)
		}
	}
}

var playEndpoint = EndpointData{
	Endpoint: play,
}

func play(ctx Context) {
	fmt.Println("play")
	query := ctx.Server.Query("q")
	core.Play(query)
	ctx.Server.SendStatus(200)
}

var saveEndpoint = EndpointData{
	Endpoint: save,
}

func save(ctx Context) {
	newConfig := new(config.Config)
	if err := ctx.Server.BodyParser(newConfig); err != nil {
		log.Fatalln("ERROR: Failed to parse new config!")
	}
	config.Save(*newConfig)
	screen.RefreshScreen()
	ctx.Server.SendStatus(200)
}

var configEndpoint = EndpointData{
	Endpoint: configuration,
}

func configuration(ctx Context) {
	configMap := config.GetMap()
	ctx.Server.JSON(configMap)
}
