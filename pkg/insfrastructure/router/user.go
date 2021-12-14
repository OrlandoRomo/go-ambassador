package router

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/middleware"
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(r *fiber.Router, c *controller.AppController) {
	user := *r
	user.Post("/users/", c.User.CreateUser)
	user.Get("/users/", middleware.AuthMiddleware, c.User.GetUser)
	user.Put("/users/", middleware.AuthMiddleware, c.User.UpdateUser)
	user.Patch("/users/password/", middleware.AuthMiddleware, c.User.UpdatePassword)
	// user.Get("/users/:user_id/links/")
}
