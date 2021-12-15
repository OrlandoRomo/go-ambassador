package repository

import (
	"strconv"

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

func (p *productRepository) CreateProduct(product *model.Product) error {
	tcx := p.db.Create(&product)
	if tcx.Error != nil {
		return tcx.Error
	}
	return nil
}

func (p *productRepository) GetProductById(id int) (*model.Product, error) {
	product := model.Product{
		ID: uint(id),
	}
	tcx := p.db.Find(&product)
	if tcx.RowsAffected == 0 {
		return nil, model.ErrNotFound{Field: "product_id", Value: strconv.Itoa(id)}
	}

	if tcx.Error != nil {
		return nil, tcx.Error
	}
	return &product, nil
}

func (p *productRepository) UpdateProductById(product *model.Product) error {
	tcx := p.db.Model(&product).Updates(&product)
	if tcx.Error != nil {
		return tcx.Error
	}
	return nil
}

func (p *productRepository) DeleteProductById(id int) error {
	tcx := p.db.Delete(&model.Product{}, id)
	if tcx.Error != nil {
		return tcx.Error
	}
	return nil
}
