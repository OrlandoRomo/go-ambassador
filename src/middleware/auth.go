package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("go_auth")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.Status(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "user unauthenticated",
		})
	}
	return c.Next()
}

func GetUserId(c *fiber.Ctx) (string, error) {
	cookie := c.Cookies("go_auth")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return "", err
	}

	payload := token.Claims.(*jwt.StandardClaims)
	return payload.Subject, nil
}
