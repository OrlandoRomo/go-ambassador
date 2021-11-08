package stat

import (
	"github.com/OrlandoRomo/go-ambassador/src/controller"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetStatAmbassador(stat fiber.Router) fiber.Router {
	stat.Get("/stats/", middleware.AuthMiddleware, controller.GetStats)
	return stat
}
