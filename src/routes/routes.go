package routes

import (
	"github.com/OrlandoRomo/go-ambassador/src/routes/auth"
	"github.com/OrlandoRomo/go-ambassador/src/routes/user"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api/v1/")
	api = auth.SetAuthRoutes(api)
	api = user.SetUserRoutes(api)
}
