package repository

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &orderRepository{db}
}

func (o *orderRepository) GetOrders() ([]*model.Order, error) {
	orders := make([]*model.Order, 0)
	tcx := o.db.Preload("OrderItems").Find(&orders)
	if tcx.Error != nil {
		return nil, tcx.Error
	}

	return orders, nil
}
