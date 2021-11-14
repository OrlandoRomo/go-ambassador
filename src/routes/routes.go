package routes

import (
	"github.com/OrlandoRomo/go-ambassador/src/routes/ambassador"
	"github.com/OrlandoRomo/go-ambassador/src/routes/auth"
	"github.com/OrlandoRomo/go-ambassador/src/routes/checkout"
	"github.com/OrlandoRomo/go-ambassador/src/routes/link"
	"github.com/OrlandoRomo/go-ambassador/src/routes/order"
	"github.com/OrlandoRomo/go-ambassador/src/routes/product"
	"github.com/OrlandoRomo/go-ambassador/src/routes/ranking"
	"github.com/OrlandoRomo/go-ambassador/src/routes/stat"
	"github.com/OrlandoRomo/go-ambassador/src/routes/user"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	admin := app.Group("/api/v1/admin/")
	admin = auth.SetAuthRoutes(admin)
	admin = user.SetUserRoutes(admin)
	admin = ambassador.SetAmbassadorAdminRoutes(admin)
	admin = product.SetProductRoutes(admin)
	admin = order.SetOrderRoutes(admin)

	ambassador := app.Group("/api/v1/ambassador/")
	ambassador = auth.SetAuthRoutes(ambassador)
	ambassador = user.SetUserAmbassadorRoutes(ambassador)
	ambassador = product.SetAmbassadorProductRoutes(ambassador)
	ambassador = link.SetLinkAmbassadorRoutes(ambassador)
	ambassador = stat.SetStatAmbassadorRoutes(ambassador)
	ambassador = ranking.SetRankingAmbassadorRoutes(ambassador)

	checkouts := app.Group("/api/v1/checkout/")
	checkouts = checkout.SetCheckoutRoutes(checkouts)
}
