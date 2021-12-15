package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	"github.com/gofiber/fiber/v2"
)

type productController struct {
	productInteractor interactor.ProductInteractor
}

type ProductController interface {
	GetProducts(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
	UpdateProductById(c *fiber.Ctx) error
	DeleteProductById(c *fiber.Ctx) error
}

func NewProductController(i interactor.ProductInteractor) ProductController {
	return &productController{i}
}

func (p *productController) GetProducts(c *fiber.Ctx) error {
	products, err := p.productInteractor.Get()
	if err != nil {
		return model.EncodeError(c, err)
	}

	return c.JSON(fiber.Map{
		"products": products,
	})
}

func (p *productController) CreateProduct(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return model.EncodeError(c, err)
	}

	if err := p.productInteractor.Create(&product); err != nil {
		return model.EncodeError(c, err)
	}
	c.Status(http.StatusCreated)
	return c.JSON(product)
}

func (p *productController) GetProductById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return model.EncodeError(c, model.ErrInvalidType{Field: "product_id"})
	}
	if id == 0 {
		return model.EncodeError(c, model.ErrInvalidType{Field: "product_id"})
	}

	product, err := p.productInteractor.GetById(id)
	if err != nil {
		return model.EncodeError(c, err)
	}
	return c.JSON(product)
}

func (p *productController) UpdateProductById(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return model.EncodeError(c, err)
	}

	id, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return model.EncodeError(c, model.ErrInvalidType{Field: "product_id"})
	}
	if id == 0 {
		return model.EncodeError(c, model.ErrInvalidType{Field: "product_id"})
	}

	_, err = p.productInteractor.GetById(id)
	if err != nil {
		return model.EncodeError(c, err)
	}

	err = p.productInteractor.Update(&product)
	if err != nil {
		return model.EncodeError(c, err)
	}

	return c.JSON(product)
}

func (p *productController) DeleteProductById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return model.EncodeError(c, model.ErrInvalidType{Field: "product_id"})
	}
	if id == 0 {
		return model.EncodeError(c, model.ErrInvalidType{Field: "product_id"})
	}

	_, err = p.productInteractor.GetById(id)
	if err != nil {
		return model.EncodeError(c, err)
	}

	err = p.productInteractor.Delete(id)
	if err != nil {
		return model.EncodeError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("product with %d was deleted successfully", id),
	})
}
