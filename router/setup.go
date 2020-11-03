package router

import (
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

// USER handles all the user routes
var USER fiber.Router

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", hello)
	USER = api.Group("/user")
	SetupUserRoutes()
}
