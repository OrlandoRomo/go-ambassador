package order

import (
	"github.com/OrlandoRomo/go-ambassador/src/controller"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetOrderRoutes(order fiber.Router) fiber.Router {
	order.Get("/orders/", middleware.AuthMiddleware, controller.GetOrders)
	return order
}
