package controller

import (
	"net/http"
	"time"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := new(model.User)

	database.DB.Where("email=?", data["email"]).First(user)

	if user.ID == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
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
		c.Status(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "user unauthenticated",
		})
	}

	token, err := middleware.GenerateJWT(user, scope)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
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

func Logout(c *fiber.Ctx) error {
	cookie := c.Cookies("go_auth")
	c.ClearCookie(cookie)
	return c.JSON(fiber.Map{
		"message": "logout successfully",
	})
}
