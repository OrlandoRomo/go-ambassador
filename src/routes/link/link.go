package link

import (
	"github.com/OrlandoRomo/go-ambassador/src/controller"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetLinkAmbassadorRoutes(link fiber.Router) fiber.Router {
	link.Post("/links/", middleware.AuthMiddleware, controller.CreateLink)
	return link
}
