package ranking

import (
	"github.com/OrlandoRomo/go-ambassador/src/controller"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetRankingAmbassadorRoutes(stat fiber.Router) fiber.Router {
	stat.Get("/rankings/", middleware.AuthMiddleware, controller.GetRankings)
	return stat
}
