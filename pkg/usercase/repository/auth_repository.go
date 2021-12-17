package repository

import "github.com/OrlandoRomo/go-ambassador/pkg/domain/model"

// Business logic for the authentication
type AuthRepository interface {
	Login(email string) (*model.User, error)
}
