package router

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, c controller.AppController) {
	admin := app.Group("/api/v1/admin/")
	SetAuthRoutes(&admin, &c)
	SetUserRoutes(&admin, &c)
	SetProductRoutes(&admin, &c)
}
