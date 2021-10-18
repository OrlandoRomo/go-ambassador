package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/middleware"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
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

	token, err := generateJWT(user)
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

func User(c *fiber.Ctx) error {
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

func Logout(c *fiber.Ctx) error {
	cookie := c.Cookies("go_auth")
	c.ClearCookie(cookie)
	return c.JSON(fiber.Map{
		"message": "logout successfully",
	})
}

func generateJWT(user *model.User) (string, error) {
	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	// TODO: replace []byte("secret") with a secured os variable
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return token, nil
}
