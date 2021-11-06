package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/gofiber/fiber/v2"
)

const (
	AscPriceSorting  = "asc"
	DescPriceSorting = "desc"
	ProductsPerPage  = 10
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

func GetProductsForFrontend(c *fiber.Ctx) error {
	var products []model.Product

	result, err := database.Cache.Get(context.Background(), "products_frontend").Result()
	if err != nil {
		tcx := database.DB.Find(&products)
		if tcx.Error != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": tcx.Error.Error(),
			})
		}
		productsMarshal, err := json.Marshal(products)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": tcx.Error.Error(),
			})
		}

		if err = database.Cache.Set(context.Background(), "products_frontend", productsMarshal, 30*time.Minute).Err(); err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"products": products,
		})
	}

	err = json.Unmarshal([]byte(result), &products)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"products": products,
	})
}

func GetProductsForBackend(c *fiber.Ctx) error {
	var products []model.Product
	searchProducts := make([]model.Product, 0)

	result, err := database.Cache.Get(context.Background(), "products_backend").Result()
	if err != nil {
		tcx := database.DB.Order("id asc").Find(&products)
		if tcx.Error != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": tcx.Error.Error(),
			})
		}
		productsMarshal, err := json.Marshal(products)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": tcx.Error.Error(),
			})
		}

		if err = database.Cache.Set(context.Background(), "products_backend", productsMarshal, 30*time.Minute).Err(); err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"products": products,
		})
	}

	err = json.Unmarshal([]byte(result), &products)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	search := c.Query("search")
	sorting := c.Query("sort")
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("the value %v is not a number", c.Query("page")),
		})
	}

	if search == "" {
		searchProducts = products
	}

	if search != "" {
		search = strings.ToLower(search)
		for _, product := range products {
			title, description := strings.ToLower(product.Title), strings.ToLower(product.Description)

			if strings.Contains(title, search) || strings.Contains(description, search) {
				searchProducts = append(searchProducts, product)
			}
		}
	}

	if sorting != "" {
		sorting = strings.ToLower(sorting)
		if sorting == AscPriceSorting {
			sort.Slice(searchProducts, func(i, j int) bool {
				return searchProducts[i].Price < searchProducts[j].Price
			})
		}
		if sorting == DescPriceSorting {
			sort.Slice(searchProducts, func(i, j int) bool {
				return searchProducts[i].Price > searchProducts[j].Price
			})
		}

	}
	total, start, end := len(searchProducts), 0, 0
	lastPage := total/ProductsPerPage + 1

	if page == 0 {
		page = 1
	}

	if page*ProductsPerPage < total {
		start = (page - 1) * ProductsPerPage
		end = start + ProductsPerPage
		searchProducts = searchProducts[start:end]
	}
	if page*ProductsPerPage > total && page <= lastPage {
		start = (page - 1) * ProductsPerPage
		end = total
		searchProducts = searchProducts[start:end]
	}

	if page > lastPage {
		searchProducts = make([]model.Product, 0)
		return c.JSON(fiber.Map{
			"products": searchProducts,
			"meta": fiber.Map{
				"total":     0,
				"page":      0,
				"last_page": 0,
			},
		})
	}

	return c.JSON(fiber.Map{
		"products": searchProducts,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage,
		},
	})
}
