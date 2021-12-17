package repository

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
	"gorm.io/gorm"
)

type linkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) repository.LinkRepository {
	return &linkRepository{db}
}

func (l *linkRepository) CreateLink(link *model.Link) error {
	tcx := l.db.Create(&link)
	if tcx.Error != nil {
		return tcx.Error
	}
	return nil
}
