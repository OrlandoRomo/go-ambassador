package registry

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	ir "github.com/OrlandoRomo/go-ambassador/pkg/interface/repository"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	ur "github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

func (r *register) NewProductController() controller.ProductController {
	return controller.NewProductController(r.NewProductInteractor())
}

func (r *register) NewProductInteractor() interactor.ProductInteractor {
	return interactor.NewProductInteractor(r.NewProductRepository())
}

func (r *register) NewProductRepository() ur.ProductRepository {
	return ir.NewProductRepository(r.db, r.redis)
}
