package checkout

import (
	"github.com/OrlandoRomo/go-ambassador/src/controller"
	"github.com/gofiber/fiber/v2"
)

func SetCheckoutRoutes(checkout fiber.Router) fiber.Router {
	checkout.Get("/links/:code", controller.GetLinksByCode)
	checkout.Post("/orders/", controller.CreateCheckoutOrders)
	checkout.Post("/orders/confirm/", nil)
	return checkout
}
