package router

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	"github.com/gofiber/fiber/v2"
)

func SetAuthRoutes(r *fiber.Router, c *controller.AppController) {
	auth := *r
	auth.Post("/login", c.Auth.Login)
	auth.Post("/logout", c.Auth.Logout)
}
