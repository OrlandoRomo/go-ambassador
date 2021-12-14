package repository

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepository{db}
}

func (p *productRepository) GetProducts() ([]*model.Product, error) {
	products := make([]*model.Product, 0)
	tcx := p.db.Find(&products)

	if tcx.Error != nil {
		return nil, tcx.Error
	}
	return products, nil
}
