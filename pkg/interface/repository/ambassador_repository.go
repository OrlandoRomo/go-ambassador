package repository

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
	"gorm.io/gorm"
)

type ambassadorRepository struct {
	db *gorm.DB
}

func NewAmbassadorRepository(db *gorm.DB) repository.AmbassadorRepository {
	return &ambassadorRepository{db}
}

func (a *ambassadorRepository) GetAmbassadors() ([]*model.User, error) {
	users := make([]*model.User, 0)
	tcx := a.db.Where("is_ambassador = true").Find(&users)
	if tcx.Error != nil {
		return nil, tcx.Error
	}
	return users, nil
}
