package dashboard

import (
	"errors"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type RadioContext struct {
	Fiber *fiber.Ctx
	Cfg   *config.SafeConfig
}

// endpointsList is a map with all endpoints.
var endpointsList = map[string]func(ctx *RadioContext){
	"config":     configuration,
	"dir":        dir,
	"stop":       stop,
	"offair":     offair,
	"yt":         yt,
	"youtubeDl":  youtubeDl,
	"playFile":   playFile,
	"playStream": playStream,
	"save":       save,
}

// findEndpoint finds endpoint by name.
func findEndpoint(name string) (func(ctx *RadioContext), error) {
	if endpointsList[name] != nil {
		return endpointsList[name], nil
	} else {
		return nil, errors.New("404")
	}
}
