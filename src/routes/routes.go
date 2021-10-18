package routes

import (
	"github.com/OrlandoRomo/go-ambassador/src/controller"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api/v1/")
	api.Post("/admin/", controller.Register)
	api.Post("/login/", controller.Login)

	admin := api.Use(middleware.AuthMiddleware)
	admin.Get("/user/", controller.User)
	admin.Post("/logout/", controller.Logout)

}
