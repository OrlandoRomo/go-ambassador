package model

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ErrInvalidType struct {
	Field string
}

func (e ErrInvalidType) Error() string {
	return fmt.Sprintf("the filed %s is not valid", e.Field)
}

type ErrInvalidCredentials struct{}

func (e ErrInvalidCredentials) Error() string {
	return "invalid credentials"
}

type ErrNotFound struct {
	Field string
	Value string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("There is not record with %s =  %s", e.Field, e.Value)
}

type ErrUnauthorized struct{}

func (e ErrUnauthorized) Error() string {
	return "user unauthorized"
}

type ErrEmailExist struct {
	Email string
}

func (e ErrEmailExist) Error() string {
	return fmt.Sprintf("The email '%s' is already used by someone else", e.Email)
}

type ErrPasswordMatch struct{}

func (e ErrPasswordMatch) Error() string {
	return "the password and confirmed don't match"
}

func EncodeError(c *fiber.Ctx, err error) error {
	// Switch for types
	switch err.(type) {
	case ErrNotFound:
		c.Status(http.StatusNotFound)
	case ErrInvalidCredentials, ErrPasswordMatch, ErrInvalidType:
		c.Status(http.StatusBadRequest)
	case ErrUnauthorized:
		c.Status(http.StatusUnauthorized)
	case ErrEmailExist:
		c.Status(http.StatusConflict)
	default:
		//Switch for error
		switch err {
		case fiber.ErrBadRequest:
			c.Status(http.StatusBadRequest)
		case gorm.ErrRecordNotFound:
			c.Status(http.StatusNotFound)
		default:
			c.Status(http.StatusInternalServerError)
		}
	}

	return c.JSON(fiber.Map{
		"error": err.Error(),
	})
}
