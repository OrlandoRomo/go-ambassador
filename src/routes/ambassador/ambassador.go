package ambassador

import (
	"github.com/OrlandoRomo/go-ambassador/src/controller"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/gofiber/fiber/v2"
)

// Create routes for admin
func SetAmbassadorAdminRoutes(ambassador fiber.Router) fiber.Router {
	ambassador.Get("/ambassadors/", middleware.AuthMiddleware, controller.GetAmbassadors)
	return ambassador
}
