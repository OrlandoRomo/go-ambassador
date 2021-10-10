package main

import (
	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Orlando")
	})

	app.Listen(":8080")
}
