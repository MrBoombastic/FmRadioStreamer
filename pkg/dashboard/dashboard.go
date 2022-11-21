package dashboard

import (
	"embed"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
)

//go:embed sites/index.html
var index embed.FS

//go:embed static/*
var static embed.FS

var app = fiber.New(fiber.Config{
	AppName:               "FmRadioStreamer",
	DisableStartupMessage: true,
})

// Init starts the dashboard
func Init(cfg *config.SafeConfig) error {
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
			handler(&RadioContext{Fiber: c, Cfg: cfg})
		}
		return nil
	})
	// Handle index file
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(index),
		PathPrefix: "sites",
	}))

	// Start!
	cfg.Lock()
	port := cfg.Port
	cfg.Unlock()
	logs.FmRadStrInfo(fmt.Sprintf("Launching dashboard at http://localhost:%v", port))
	return app.Listen(fmt.Sprintf(":%v", port))
}
