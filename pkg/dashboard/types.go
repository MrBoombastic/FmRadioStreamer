package dashboard

import "github.com/gofiber/fiber/v2"

type Context struct {
	Server fiber.Ctx
}

type EndpointData struct {
	Endpoint func(ctx Context)
}

func NewContext(c *fiber.Ctx) Context {
	return Context{
		Server: *c,
	}
}
