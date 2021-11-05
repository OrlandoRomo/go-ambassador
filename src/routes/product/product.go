package product

import (
	"github.com/OrlandoRomo/go-ambassador/src/controller"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetProductRoutes(product fiber.Router) fiber.Router {
	product.Get("/products/", middleware.AuthMiddleware, controller.GetProducts)
	product.Post("/products/", middleware.AuthMiddleware, controller.CreateProduct)
	product.Get("/products/:product_id", middleware.AuthMiddleware, controller.GetProductById)
	product.Put("/products/:product_id", middleware.AuthMiddleware, controller.UpdateProductById)
	product.Delete("/products/:product_id", middleware.AuthMiddleware, controller.DeleteProductById)
	return product
}

func SetAmbassadorProductRoutes(product fiber.Router) fiber.Router {
	product.Get("/products/frontend/", controller.GetProductsForFrontend)
	return product
}
