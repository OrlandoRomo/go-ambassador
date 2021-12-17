package router

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/middleware"
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	"github.com/gofiber/fiber/v2"
)

func SetLinkRoutes(r *fiber.Router, c *controller.AppController) {
	link := *r
	link.Post("/links/", middleware.AuthMiddleware, c.Link.CreateLink)
}
