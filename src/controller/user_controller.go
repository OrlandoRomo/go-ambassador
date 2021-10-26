package controller

import (
	"net/http"
	"strconv"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/gofiber/fiber/v2"
)

func NewUser(c *fiber.Ctx) error {
	var body map[string]string
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "bad request",
		})
	}

	if body["password"] != body["confirmed_password"] {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "password does not match",
		})
	}

	admin := model.NewUser(body)

	admin.SetPassword(body["password"])

	database.DB.Create(&admin)

	c.Status(http.StatusCreated)
	return c.JSON(&admin)
}

func GetUser(c *fiber.Ctx) error {
	idUser, err := middleware.GetUserId(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "error getting user",
		})
	}
	var user model.User
	database.DB.Where("id = ?", idUser).Find(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	var body map[string]string
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "bad request",
		})
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	user := model.User{
		ID:        uint(id),
		FirstName: body["first_name"],
		LastName:  body["last_name"],
		Email:     body["email"],
	}
	tx := database.DB.Model(&user).Updates(&user)
	if tx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tx.Error.Error(),
		})
	}

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var body map[string]string
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "bad request",
		})
	}

	if body["password"] != body["confirmed_password"] {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "password does not match",
		})
	}

	userId, err := middleware.GetUserId(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := model.User{
		ID: uint(id),
	}

	user.SetPassword(body["password"])

	tx := database.DB.Model(&user).Updates(&user)
	if tx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tx.Error.Error(),
		})
	}
	c.Status(http.StatusNoContent)
	return nil

}
