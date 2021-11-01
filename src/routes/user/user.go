package user

import (
	"github.com/OrlandoRomo/go-ambassador/src/controller"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(user fiber.Router) fiber.Router {
	user.Post("/users/", controller.CreateUser)
	user.Get("/users/", middleware.AuthMiddleware, controller.GetUser)
	user.Put("/users/", middleware.AuthMiddleware, controller.UpdateUser)
	user.Put("/users/password/", middleware.AuthMiddleware, controller.UpdatePassword)
	return user
}
