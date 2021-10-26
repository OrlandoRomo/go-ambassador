package auth

import (
	"github.com/OrlandoRomo/go-ambassador/src/controller"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetAuthRoutes(auth fiber.Router) fiber.Router {
	auth.Post("/login/", controller.Login)
	auth.Post("/logout/", middleware.AuthMiddleware, controller.Logout)
	return auth
}
