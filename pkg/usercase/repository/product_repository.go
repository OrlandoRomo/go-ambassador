package repository

import "github.com/OrlandoRomo/go-ambassador/pkg/domain/model"

type ProductRepository interface {
	GetProducts() ([]*model.Product, error)
	CreateProduct(*model.Product) error
	GetProductById(id int) (*model.Product, error)
	UpdateProductById(*model.Product) error
	DeleteProductById(id int) error
}
