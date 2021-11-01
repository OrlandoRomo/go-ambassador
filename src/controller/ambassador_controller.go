package controller

import (
	"net/http"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/gofiber/fiber/v2"
)

func GetAmbassadors(c *fiber.Ctx) error {
	var users []model.User
	tcx := database.DB.Where("is_ambassador = true;").Find(&users)
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"ambassadors": users,
	})
}
