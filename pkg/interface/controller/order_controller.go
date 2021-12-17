package controller

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	"github.com/gofiber/fiber/v2"
)

type orderController struct {
	orderInteractor interactor.OrderInteractor
}

type OrderController interface {
	GetOrders(c *fiber.Ctx) error
}

func NewOrderController(i interactor.OrderInteractor) OrderController {
	return &orderController{i}
}

func (o *orderController) GetOrders(c *fiber.Ctx) error {
	orders, err := o.orderInteractor.Get()
	if err != nil {
		return model.EncodeError(c, err)
	}

	return c.JSON(fiber.Map{
		"orders": orders,
	})
}
