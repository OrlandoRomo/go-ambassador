package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	AscPriceSorting  = "asc"
	DescPriceSorting = "desc"
	ProductsPerPage  = 10
)

type productRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewProductRepository(db *gorm.DB, redis *redis.Client) repository.ProductRepository {
	return &productRepository{db, redis}
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

func (p *productRepository) GetProductsBackend(s *model.SearchProduct) (interface{}, error) {
	result, err := p.redis.Get(context.Background(), "products_backend").Result()
	if err != nil {
		tcx := p.db.Order("id asc").Find(&s.Result)
		if tcx.Error != nil {
			return nil, tcx.Error
		}
		productsMarshal, err := json.Marshal(s.Result)
		if err != nil {
			return nil, err
		}
		statusCmd := p.redis.Set(context.Background(), "products_backend", productsMarshal, 0)
		fmt.Println(statusCmd.Err())
		fmt.Println(statusCmd.String())
		fmt.Println(statusCmd.Name())
		fmt.Println(statusCmd.Result())
		// if err = p.redis.Set(context.Background(), "products_backend", productsMarshal, 30*time.Minute).Err(); err != nil {
		// 	return nil, err
		// }

		searchByCriteria(s)

		return setPager(s)
	}

	if err := json.Unmarshal([]byte(result), &s.Result); err != nil {
		return nil, err
	}

	searchByCriteria(s)

	return setPager(s)
}

func (p *productRepository) GetProductsFrontend(s *model.SearchProduct) (interface{}, error) {
	return nil, nil
}

func searchByCriteria(s *model.SearchProduct) {
	searchedProducts := make([]*model.Product, len(s.Result))
	copy(searchedProducts, s.Result)

	if s.Search != "" {
		search := strings.ToLower(s.Search)
		for index, product := range searchedProducts {
			title, description := strings.ToLower(product.Title), strings.ToLower(product.Description)
			if !strings.Contains(title, search) || !strings.Contains(description, search) {
				s.Result = append(s.Result[:index], s.Result[index+1:]...)
			}
		}
	}

	if s.Sort != "" {
		sorting := strings.ToLower(s.Sort)
		if sorting == AscPriceSorting {
			sort.Slice(s.Result, func(i, j int) bool {
				return s.Result[i].Price > s.Result[j].Price
			})
		}

		if sorting == DescPriceSorting {
			sort.Slice(s.Result, func(i, j int) bool {
				return s.Result[i].Price < s.Result[j].Price
			})
		}
	}
}

func setPager(s *model.SearchProduct) (interface{}, error) {

	total, start, end := len(s.Result), 0, 0
	lastPage := total / ProductsPerPage
	if lastPage == 0 {
		lastPage++
	}

	if s.Page == 0 {
		s.Page = 1
	}

	if s.Page*ProductsPerPage < total {
		start = (s.Page - 1) * ProductsPerPage
		end = start + ProductsPerPage
		s.Result = s.Result[start:end]
	}
	if s.Page*ProductsPerPage > total && s.Page <= lastPage {
		start = (s.Page - 1) * ProductsPerPage
		end = total
		s.Result = s.Result[start:end]
	}

	if s.Page > lastPage {
		return nil, model.ErrPage{s.Page}
	}
	return fiber.Map{
		"products": s.Result,
		"meta": fiber.Map{
			"total":     total,
			"page":      s.Page,
			"last_page": lastPage,
		},
	}, nil
}
