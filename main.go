package main

import (
	"flag"
	"os"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stripe/stripe-go/v72"
)

func main() {
	var StripeSecretApiKey = flag.String("stripe_secret_api_key", envString("STRIPE_SECRET_API_KEY", ""), "Set up stripe secret api key")
	flag.Parse()
	database.Connect()
	database.AutoMigrate()
	database.SetUpRedis()

	stripe.Key = *StripeSecretApiKey

	app := fiber.New()

	// Enabling CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetRoutes(app)

	app.Listen(":8000")
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
