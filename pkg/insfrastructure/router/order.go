package router

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/middleware"
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	"github.com/gofiber/fiber/v2"
)

func SetOrderRoutes(r *fiber.Router, c *controller.AppController) {
	order := *r
	order.Get("/orders/", middleware.AuthMiddleware, c.Order.GetOrders)
}
