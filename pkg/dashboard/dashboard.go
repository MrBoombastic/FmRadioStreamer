package dashboard

import (
	"embed"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"io/ioutil"
	"net/http"
)

//go:embed sites/index.html
var index embed.FS

//go:embed static/*
var static embed.FS

// musicList returns slice of all files in 'music' directory
func musicList() ([]string, error) {
	files, err := ioutil.ReadDir("music/")
	if err != nil {
		return nil, err
	}
	var filesSlice []string
	for _, item := range files {
		filesSlice = append(filesSlice, item.Name())
	}
	return filesSlice, nil
}

var app = fiber.New()

// Init starts the dashboard
func Init() {
	// Handle static files
	app.Use("/static/", filesystem.New(filesystem.Config{
		Root:       http.FS(static),
		PathPrefix: "static",
	}))
	// Handle API endpoints
	app.Use("/api/*", func(c *fiber.Ctx) error {
		endpoint := fmt.Sprintf("%s", c.Params("*"))
		foundEndpoint, err := findEndpoint(endpoint)
		if err != nil {
			return c.SendStatus(404)
		}
		handler := foundEndpoint
		if handler != nil {
			handler(c)
		}
		return nil
	})
	// Handle index file
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(index),
		PathPrefix: "sites",
	}))

	// Start!
	app.Listen(fmt.Sprintf(":%v", config.GetPort()))
}
