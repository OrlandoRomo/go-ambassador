package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/cache"
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/db"
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/router"
	"github.com/OrlandoRomo/go-ambassador/pkg/registry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	flag.Parse()
	config, err := db.NewConfig()
	if err != nil {
		log.Println(
			"message", "could not reach out the app configuration",
			"severity", "ERR")
		return
	}

	ambassadorDB, err := db.NewDB(&config.DB)
	if err != nil {
		log.Println("ERR", err.Error())
		return
	}

	err = db.AutoMigrate(ambassadorDB)
	if err != nil {
		log.Println(
			"message", err.Error(),
			"severity", "ERR",
		)
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

	listenAndServe(app)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func listenAndServe(app *fiber.App) {
	connClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT)
		<-sigint
		log.Println(
			"message", "stopping go-ambassador",
			"severity", "INFO",
		)
		if err := app.Shutdown(); err != nil {
			log.Println(err)
		}
		close(connClosed)
	}()

	if err := app.Listen(":8000"); err != http.ErrServerClosed {
		log.Println(
			"message", err.Error(),
			"severity", "CRITICAL",
		)
	} else {
		log.Println(
			"message", "service stopped",
			"severity", "NOTICE",
		)
	}
	<-connClosed

}
