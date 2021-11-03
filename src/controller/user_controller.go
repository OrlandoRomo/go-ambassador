package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUser(c *fiber.Ctx) error {
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

	admin := &model.User{
		FirstName:    body["first_name"],
		LastName:     body["last_name"],
		Email:        body["email"],
		IsAmbassador: middleware.IsAmbassadorPath(c),
	}

	admin.SetPassword(body["password"])

	tcx := database.DB.Create(&admin)
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

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
	tcx := database.DB.Where("id = ?", idUser).Find(&user)
	if tcx.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("there is not a user with id `%s`", idUser),
		})
	}
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	if middleware.IsAmbassadorPath(c) {
		ambassador := model.Ambassador(user)
		revenue, err := calculateRevenue(c, ambassador.ID)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		ambassador.Revenue = revenue
		return c.JSON(ambassador)
	}
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
	tcx := database.DB.Model(&user).Updates(&user)
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
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

	tcx := database.DB.Model(&user).Updates(&user)
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}
	c.Status(http.StatusNoContent)
	return nil
}

func calculateRevenue(c *fiber.Ctx, ID uint) (float64, error) {
	var orders []model.Order
	var revenue float64
	tcx := database.DB.Preload("OrderItems").Find(&orders, &model.Order{
		UserID:      ID,
		IsCompleted: true,
	})
	if tcx.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return 0.0, gorm.ErrRecordNotFound
	}
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return 0.0, tcx.Error
	}

	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.AdminRevenue
		}
	}

	return revenue, nil
}
