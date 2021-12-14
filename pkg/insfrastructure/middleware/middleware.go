package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

const (
	SecretKey  = "secret"
	Admin      = "admin"
	Ambassador = "ambassador"
)

type ClaimsScope struct {
	jwt.StandardClaims
	Scope string
}

func AuthMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("go_auth")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsScope{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		c.Status(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "user unauthenticated",
		})
	}

	payload := token.Claims.(*ClaimsScope)

	if (payload.Scope == Admin && IsAmbassadorPath(c)) || (payload.Scope == Ambassador && !IsAmbassadorPath(c)) {
		c.Status(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "user unauthenticated",
		})
	}

	return c.Next()
}

func IsAmbassadorPath(c *fiber.Ctx) bool {
	return strings.Contains(c.Path(), "/api/v1/ambassador/")
}

func GetUserId(c *fiber.Ctx) (int, error) {
	cookie := c.Cookies("go_auth")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsScope{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*ClaimsScope)
	id, err := strconv.Atoi(payload.Subject)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GenerateJWT(user model.User, scope string) (string, error) {
	payload := ClaimsScope{
		Scope: scope,
	}
	payload.Subject = strconv.Itoa(int(user.ID))
	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

	// TODO: replace []byte("secret") with a secured os variable
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
