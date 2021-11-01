package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []model.Product

	tcx := database.DB.Find(&products)
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"products": products,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tcx := database.DB.Create(&product)
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}
	c.Status(http.StatusCreated)
	return c.JSON(product)
}

func GetProductById(c *fiber.Ctx) error {
	var product model.Product
	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	product.ID = uint(productID)

	tcx := database.DB.Find(&product)
	if tcx.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("there is not a product with id `%d`", productID),
		})
	}

	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	return c.JSON(product)
}

func UpdateProductById(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	product.ID = uint(productID)

	tcx := database.DB.First(&model.Product{}, productID)
	if tcx.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("there is not a product with id `%d`", productID),
		})
	}

	tcx = database.DB.Model(&product).Updates(&product)
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	return c.JSON(product)
}

func DeleteProductById(c *fiber.Ctx) error {
	tcx := database.DB.Delete(&model.Product{}, c.Params("product_id"))
	if tcx.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("there is not a product with id `%s`", c.Params("product_id")),
		})
	}
	c.Status(http.StatusNoContent)
	return nil
}
