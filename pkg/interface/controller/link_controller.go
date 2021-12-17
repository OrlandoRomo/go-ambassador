package controller

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/middleware"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type linkController struct {
	linkInteractor interactor.LinkInteractor
}

type LinkController interface {
	CreateLink(c *fiber.Ctx) error
}

func NewLinkController(i interactor.LinkInteractor) LinkController {
	return &linkController{i}
}

func (a *linkController) CreateLink(c *fiber.Ctx) error {
	var request model.LinkRequest
	if err := c.BodyParser(&request); err != nil {
		return model.EncodeError(c, err)
	}
	id, err := middleware.GetUserId(c)
	if err != nil {
		return model.EncodeError(c, err)
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

	if err := a.linkInteractor.Create(&link); err != nil {
		return model.EncodeError(c, err)
	}

	return c.JSON(link)
}
