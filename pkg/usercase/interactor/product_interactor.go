package interactor

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

type productInteractor struct {
	productRepository repository.ProductRepository
}

type ProductInteractor interface {
	Get() ([]*model.Product, error)
}

func NewProductInteractor(r repository.ProductRepository) ProductInteractor {
	return &productInteractor{r}
}

func (p *productInteractor) Get() ([]*model.Product, error) {
	products, err := p.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
