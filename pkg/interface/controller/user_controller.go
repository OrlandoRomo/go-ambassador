package controller

import (
	"net/http"
	"strconv"

	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/middleware"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	"github.com/gofiber/fiber/v2"
)

const (
	SearchByEmail = "email"
	SearchById    = "id"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	CreateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	UpdatePassword(c *fiber.Ctx) error
}

func NewUserController(i interactor.UserInteractor) UserController {
	return &userController{i}
}

func (u *userController) CreateUser(c *fiber.Ctx) error {
	var data map[string]string
	admin := new(model.User)
	if err := c.BodyParser(&data); err != nil {
		return model.EncodeError(c, err)
	}

	admin = &model.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		IsAmbassador: middleware.IsAmbassadorPath(c),
	}

	// Verify whether a user already exists with data["email"] or not
	if err := u.userInteractor.Get(admin, SearchByEmail); err != nil {
		return model.EncodeError(c, err)
	}

	if admin != nil && admin.ID != 0 {
		return model.EncodeError(c, model.ErrEmailExist{Email: admin.Email})
	}

	admin.SetPassword(data["password"])

	if err := u.userInteractor.Create(admin); err != nil {
		return model.EncodeError(c, err)
	}
	c.Status(http.StatusCreated)
	return c.JSON(&admin)
}

func (u *userController) GetUser(c *fiber.Ctx) error {
	id, err := middleware.GetUserId(c)
	if err != nil {
		return model.EncodeError(c, err)
	}

	user := model.User{
		ID: uint(id),
	}
	if err = u.userInteractor.Get(&user, SearchById); err != nil {
		return model.EncodeError(c, err)
	}

	if user.ID == 0 {
		return model.EncodeError(c, model.ErrNotFound{Field: "ID", Value: strconv.Itoa(id)})
	}

	// if middleware.IsAmbassadorPath(c) {
	// 	// ambassador := model.Ambassador(user)
	// 	// return, err:= u
	// }

	return c.JSON(user)
}

func (u *userController) UpdateUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return model.EncodeError(c, err)
	}
	idUser, err := middleware.GetUserId(c)
	if err != nil {
		return model.EncodeError(c, err)
	}

	user := model.User{
		ID:        uint(idUser),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	if err = u.userInteractor.Update(&user); err != nil {
		return model.EncodeError(c, err)
	}
	return c.JSON(&user)
}

func (u *userController) UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return model.EncodeError(c, err)
	}

	if data["password"] != data["confirmed_password"] {
		return model.EncodeError(c, model.ErrPasswordMatch{})
	}

	idUser, err := middleware.GetUserId(c)
	if err != nil {
		return model.EncodeError(c, err)
	}

	user := model.User{
		ID: uint(idUser),
	}

	user.SetPassword(data["password"])

	if err = u.userInteractor.Update(&user); err != nil {
		return model.EncodeError(c, err)
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
