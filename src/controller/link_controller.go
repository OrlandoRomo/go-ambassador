package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/gofiber/fiber/v2"
)

func GetLinksByUserId(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	var links []model.Link
	tcx := database.DB.Preload("Products").Where("user_id = ?", userId).Find(&links)
	if tcx.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("there are not links with user_id `%s`", c.Params("user_id")),
		})
	}
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	for i, link := range links {
		var orders []model.Order
		tcx = database.DB.Where("code = ? AND is_completed = true", link.Code).Find(&orders)
		if tcx.Error != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": tcx.Error.Error(),
			})
		}
		links[i].Orders = orders

	}

	return c.JSON(fiber.Map{
		"links": links,
	})
}
