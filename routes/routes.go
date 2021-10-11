package routes

import (
	"github.com/OrlandoRomo/go-ambassador/controller"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api/v1/")
	api.Post("/admin/", controller.Register)
	api.Post("/login/", controller.Login)
}
