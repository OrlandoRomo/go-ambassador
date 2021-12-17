package repository

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(us *model.User) error {
	tcx := u.db.Create(&us)
	if tcx.Error != nil {
		return tcx.Error
	}
	return nil
}

func (u *userRepository) GetUser(us *model.User, searchType string) error {
	if searchType == controller.SearchByEmail {
		tcx := u.db.Where("email = ?", us.Email).First(&us)
		if tcx.Error != nil {
			return tcx.Error
		}
	}

	if searchType == controller.SearchById {
		tcx := u.db.Where("id = ?", us.ID).First(&us)
		if tcx.Error != nil {
			return tcx.Error
		}
	}
	return nil
}

func (u *userRepository) UpdateUser(us *model.User) error {
	tcx := u.db.Model(&us).Updates(&us)
	if tcx.Error != nil {
		return tcx.Error
	}
	return nil
}
