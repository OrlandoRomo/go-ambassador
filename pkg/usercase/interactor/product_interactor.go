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
	Create(*model.Product) error
	GetById(id int) (*model.Product, error)
	Update(*model.Product) error
	Delete(id int) error
	Cache(*model.SearchProduct) (interface{}, error)
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

func (p *productInteractor) Create(product *model.Product) error {
	err := p.productRepository.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (p *productInteractor) GetById(id int) (*model.Product, error) {
	product, err := p.productRepository.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productInteractor) Update(product *model.Product) error {
	if err := p.productRepository.UpdateProductById(product); err != nil {
		return err
	}
	return nil
}

func (p *productInteractor) Delete(id int) error {
	if err := p.productRepository.DeleteProductById(id); err != nil {
		return err
	}
	return nil
}

func (p *productInteractor) Cache(s *model.SearchProduct) (interface{}, error) {
	result, err := p.productRepository.GetProductsBackend(s)
	if err != nil {
		return nil, err
	}
	return result, nil
}
