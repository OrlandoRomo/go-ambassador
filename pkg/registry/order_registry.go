package registry

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	ir "github.com/OrlandoRomo/go-ambassador/pkg/interface/repository"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	ur "github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

func (r *register) NewOrderController() controller.OrderController {
	return controller.NewOrderController(r.NewOrderInteractor())
}

func (r *register) NewOrderInteractor() interactor.OrderInteractor {
	return interactor.NewOrderInteractor(r.NewOrderRepository())
}

func (r *register) NewOrderRepository() ur.OrderRepository {
	return ir.NewOrderRepository(r.db)
}
