package controller

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	"github.com/gofiber/fiber/v2"
)

type productController struct {
	productInteractor interactor.ProductInteractor
}

type ProductController interface {
	GetProducts(c *fiber.Ctx) error
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
