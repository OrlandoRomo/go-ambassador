package interactor

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

type orderInteractor struct {
	orderRepository repository.OrderRepository
}

type OrderInteractor interface {
	Get() ([]*model.Order, error)
}

func NewOrderInteractor(r repository.OrderRepository) OrderInteractor {
	return &orderInteractor{r}
}

func (o *orderInteractor) Get() ([]*model.Order, error) {
	orders, err := o.orderRepository.GetOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
