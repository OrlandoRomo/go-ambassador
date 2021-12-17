package router

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/middleware"
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	"github.com/gofiber/fiber/v2"
)

func SetAmbassadorRoutes(r *fiber.Router, c *controller.AppController) {
	ambassador := *r
	ambassador.Get("/ambassadors/", middleware.AuthMiddleware, c.Ambassador.GetAmbassadors)
}
