package dashboard

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

// endpointsList is a map with all endpoints
var endpointsList = map[string]func(ctx *fiber.Ctx){
	"config":     configuration,
	"music":      music,
	"stop":       stop,
	"offair":     offair,
	"yt":         yt,
	"youtubeDl":  youtubeDl,
	"playFile":   playFile,
	"playStream": playStream,
	"save":       save,
}

// findEndpoint finds endpoint by name
func findEndpoint(name string) (func(ctx *fiber.Ctx), error) {
	if endpointsList[name] != nil {
		return endpointsList[name], nil
	} else {
		return nil, errors.New("404")
	}
}
