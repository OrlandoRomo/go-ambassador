package repository

import "github.com/OrlandoRomo/go-ambassador/pkg/domain/model"

type UserRepository interface {
	CreateUser(*model.User) error
	GetUser(*model.User, string) error
	UpdateUser(*model.User) error
}
