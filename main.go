package main

import (
	"github.com/OrlandoRomo/go-ambassador/routes"
	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	// Enabling CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetRoutes(app)

	app.Listen(":8080")
}
