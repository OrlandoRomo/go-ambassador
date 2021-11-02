package controller

import (
	"net/http"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
	var orders []model.Order
	tcx := database.DB.Preload("OrderItems").Find(&orders)
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	for i, order := range orders {
		orders[i].Total = order.GetTotal()
	}

	return c.JSON(fiber.Map{
		"orders": orders,
	})
}
