package main

import (
	"flag"
	"os"

	"github.com/OrlandoRomo/go-ambassador/src/controller"
	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stripe/stripe-go/v72"
)

func main() {
	var (
		StripeSecretApiKey = flag.String("stripe_secret_api_key", envString("STRIPE_SECRET_API_KEY", ""), "Set up stripe secret api key")
		MailHogHost        = flag.String("mail_hog_host", envString("MAIL_HOG_HOST", ""), "Mailhog host")
		MailHostFrom       = flag.String("mail_hog_from", envString("MAIL_HOG_FROM", ""), "Mailhog email from")
		MailHogAdminEmail  = flag.String("mail_hog_admin_email", envString("MAIL_HOG_ADMIN_EMAIL", ""), "Mailhog admin email")
	)

	flag.Parse()
	database.Connect()
	database.AutoMigrate()
	database.SetUpRedis()

	// Refactor these dependencies :p
	stripe.Key = *StripeSecretApiKey
	controller.MailHogClient.Host = *MailHogHost
	controller.MailHogClient.From = *MailHostFrom
	controller.MailHogClient.AdminEmail = *MailHogAdminEmail

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
