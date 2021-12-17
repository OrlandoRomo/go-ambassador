package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/middleware"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authInteractor interactor.AuthInteractor
}

type AuthController interface {
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

func NewAuthController(a interactor.AuthInteractor) AuthController {
	return &authController{a}
}

func (a *authController) Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println(err)
		return model.EncodeError(c, err)
	}
	user, err := a.authInteractor.Login(data["email"])
	if err != nil {
		return model.EncodeError(c, err)
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		return model.EncodeError(c, err)
	}

	isAmbassador := middleware.IsAmbassadorPath(c)
	scope := ""
	if isAmbassador {
		scope = middleware.Ambassador
	}
	if !isAmbassador {
		scope = middleware.Admin
	}

	if !isAmbassador && user.IsAmbassador {
		return model.EncodeError(c, model.ErrUnauthorized{})
	}

	token, err := middleware.GenerateJWT(*user, scope)
	if err != nil {
		return model.EncodeError(c, err)
	}

	cookie := fiber.Cookie{
		Name:     "go_auth",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	c.Status(http.StatusOK)
	return c.JSON(fiber.Map{
		"message": "log in success",
	})
}

func (a *authController) Logout(c *fiber.Ctx) error {
	cookie := c.Cookies("go_auth")
	c.ClearCookie(cookie)
	return c.JSON(fiber.Map{
		"message": "logout successfully",
	})
}
