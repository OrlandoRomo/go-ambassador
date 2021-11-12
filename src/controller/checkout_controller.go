package controller

import (
	"fmt"
	"net/http"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/gofiber/fiber/v2"
)

func GetLinksByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("the value %s is not valid", code),
		})
	}

	link := model.Link{
		Code: code,
	}

	tcx := database.DB.Preload("User").Preload("Products").First(&link)
	if tcx.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("there is not a link with code %s\n", code),
		})
	}
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	return c.JSON(link)
}

func CreateCheckoutOrders(c *fiber.Ctx) error {
	var request model.OrderRequest
	if err := c.BodyParser(&request); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	link := model.Link{
		Code: request.Code,
	}
	tcx := database.DB.Preload("User").First(&link)
	if tcx.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("there is not a link with code %s\n", request.Code),
		})
	}
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}
	order := model.Order{
		Code:            link.Code,
		UserID:          link.UserID,
		AmbassadorEmail: link.User.Email,
		FirstName:       request.FirstName,
		LastName:        request.LastName,
		Email:           request.Email,
		Address:         request.Address,
		City:            request.Country,
		Zip:             request.Zip,
	}
	tcx = database.DB.Begin()

	if err := tcx.Create(&order).Error; err != nil {
		tcx.Rollback()
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	for _, p := range request.Products {
		product := model.Product{
			ID: uint(p["product_id"]),
		}

		tcx = database.DB.First(&product)
		if tcx.RowsAffected == 0 {
			c.Status(http.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": fmt.Sprintf("there is not a prodict with %d\n", p["product_id"]),
			})
		}
		if tcx.Error != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": tcx.Error.Error(),
			})
		}

		total := product.Price * float64(p["quantity"])
		item := model.OrderItem{
			OrderID:           order.ID,
			ProductTitle:      product.Title,
			Price:             product.Price,
			Quantity:          uint(p["quantity"]),
			AmbassadorRevenue: 0.1 * total,
			AdminRevenue:      0.9 * total,
		}

		if err := tcx.Create(&item).Error; err != nil {
			tcx.Rollback()
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": tcx.Error.Error(),
			})
		}
	}

	tcx.Commit()

	return c.JSON(order)
}

func ConfirmCheckoutOrders(c *fiber.Ctx) error {
	return nil
}
