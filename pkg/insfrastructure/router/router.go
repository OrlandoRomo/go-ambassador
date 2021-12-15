package router

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	"github.com/gofiber/fiber/v2"
)

const (
	AdminVersion      = "/api/v2/admin/"
	AmbassadorVersion = "/api/v2/ambassador/"
)

func NewRouter(app *fiber.App, c controller.AppController) {
	admin := app.Group(AdminVersion)
	SetAuthRoutes(&admin, &c)
	SetUserRoutes(&admin, &c, AdminVersion)
	SetProductRoutes(&admin, &c)

	ambassador := app.Group(AmbassadorVersion)
	SetAuthRoutes(&ambassador, &c)
	SetUserRoutes(&ambassador, &c, AmbassadorVersion)

}
