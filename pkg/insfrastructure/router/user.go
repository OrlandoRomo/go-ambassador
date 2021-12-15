package router

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/middleware"
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(r *fiber.Router, c *controller.AppController, version string) {
	user := *r
	user.Post("/users/", c.User.CreateUser)
	user.Get("/users/", middleware.AuthMiddleware, c.User.GetUser)
	user.Put("/users/", middleware.AuthMiddleware, c.User.UpdateUser)
	user.Patch("/users/password/", middleware.AuthMiddleware, c.User.UpdatePassword)
	if version == AdminVersion {
		// user.Get("/users/:user_id/links/") this for /api/v1/admin/
	}
	if version == AmbassadorVersion {
		// user.Get("/users/rankings/", middleware.AuthMiddleware, controller.GetRankings) this is for /api/v1/ambassador/
	}
}
