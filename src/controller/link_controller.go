package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func CreateLink(c *fiber.Ctx) error {
	var request model.LinkRequest
	if err := c.BodyParser(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		c.Status(http.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	link := model.Link{
		UserID: uint(id),
		Code:   uuid.NewString(),
	}
	for _, productId := range request.Products {
		product := model.Product{
			ID: uint(productId),
		}
		link.Products = append(link.Products, product)
	}
	tcx := database.DB.Create(&link)
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(link)
}

func GetStats(c *fiber.Ctx) error {
	stats := make([]model.Stat, 0)
	links := make([]model.Link, 0)
	orders := make([]model.Order, 0)

	userId, err := middleware.GetUserId(c)
	if err != nil {
		c.Status(http.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	tcx := database.DB.Find(&links, model.Link{
		UserID: uint(id),
	})
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	for _, link := range links {
		tcx = database.DB.Preload("OrderItems").Find(&orders, model.Order{
			Code:        link.Code,
			IsCompleted: true,
		})
		if tcx.RowsAffected == 0 {
			continue
		}
		if tcx.Error != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		revenue := 0.0
		for _, order := range orders {
			revenue += order.GetTotal()
		}

		stats = append(stats, model.Stat{
			Code:    link.Code,
			Count:   len(orders),
			Revenue: revenue,
		})
	}

	return c.JSON(fiber.Map{
		"stats": stats,
	})
}
