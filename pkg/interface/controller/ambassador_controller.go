package controller

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	"github.com/gofiber/fiber/v2"
)

type ambassadorController struct {
	ambassadorInteractor interactor.AmbassadorInteractor
}

type AmbassadorController interface {
	GetAmbassadors(c *fiber.Ctx) error
}

func NewAmbassadorController(i interactor.AmbassadorInteractor) AmbassadorController {
	return &ambassadorController{i}
}

func (a *ambassadorController) GetAmbassadors(c *fiber.Ctx) error {
	ambassadors, err := a.ambassadorInteractor.Get()
	if err != nil {
		return model.EncodeError(c, err)
	}

	return c.JSON(fiber.Map{
		"ambassadors": ambassadors,
	})
}
