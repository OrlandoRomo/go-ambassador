package repository

import "github.com/OrlandoRomo/go-ambassador/pkg/domain/model"

type OrderRepository interface {
	GetOrders() ([]model.Order, error)
}
