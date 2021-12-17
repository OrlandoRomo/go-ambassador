package repository

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repository.AuthRepository {
	return &authRepository{db}
}

func (a *authRepository) Login(email string) (*model.User, error) {
	user := new(model.User)
	tcx := a.db.Where("email=?", email).First(user)

	if tcx.Error != nil {
		return nil, tcx.Error
	}

	return user, nil
}
