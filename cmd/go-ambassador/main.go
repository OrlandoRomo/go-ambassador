package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/cache"
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/db"
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/router"
	"github.com/OrlandoRomo/go-ambassador/pkg/registry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	var (
	// StripeSecretApiKey = flag.String("stripe_secret_api_key", envString("STRIPE_SECRET_API_KEY", ""), "Set up stripe secret api key")
	// MailHogHost        = flag.String("mail_hog_host", envString("MAIL_HOG_HOST", ""), "Mailhog host")
	// MailHostFrom       = flag.String("mail_hog_from", envString("MAIL_HOG_FROM", ""), "Mailhog email from")
	// MailHogAdminEmail  = flag.String("mail_hog_admin_email", envString("MAIL_HOG_ADMIN_EMAIL", ""), "Mailhog admin email")
	// ConfigFile = flag.String("db_config_file", envString("DB_CONFIG_FILE", ""), "DB configuration file")
	)

	flag.Parse()
	config, err := db.NewConfig()
	if err != nil {
		log.Println("ERR", "could not reach out the app configuration")
		return
	}

	ambassadorDB, err := db.NewDB(&config.DB)
	if err != nil {
		log.Println("ERR", err.Error())
		return
	}

	err = db.AutoMigrate(ambassadorDB)
	if err != nil {
		log.Println("ERR", err.Error())
		return
	}
	redis := cache.NewCache(config.Redis.Port)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowMethods:     fmt.Sprintf("%s,%s,%s,%s,%s,%s", http.MethodGet, http.MethodPost, http.MethodHead, http.MethodPut, http.MethodDelete, http.MethodPatch),
		AllowCredentials: true,
	}))

	r := registry.NewRegister(ambassadorDB, redis)
	router.NewRouter(app, r.NewAppController())

	app.Listen(":8000")

	// database.Connect()
	// database.AutoMigrate()
	// database.SetUpRedis()

	// // Refactor these dependencies :p
	// stripe.Key = *StripeSecretApiKey
	// controller.MailHogClient.Host = *MailHogHost
	// controller.MailHogClient.From = *MailHostFrom
	// controller.MailHogClient.AdminEmail = *MailHogAdminEmail

	// app := fiber.New()

	// // Enabling CORS
	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	// routes.SetRoutes(app)

	// app.Listen(":8000")
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
