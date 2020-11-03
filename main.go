package main

import (
	"log"

	"github.com/hooneun/go-postgres-auth/database"
	"github.com/hooneun/go-postgres-auth/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CreateServer creawtes a new Fiber instance
func CreateServer() *fiber.App {
	app := fiber.New()

	return app
}

func main() {
	database.ConnectToDB()
	app := CreateServer()

	app.Use(cors.New())

	router.SetupRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3333"))
}
