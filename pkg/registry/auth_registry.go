package registry

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	ir "github.com/OrlandoRomo/go-ambassador/pkg/interface/repository"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	ur "github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

func (r *register) NewAuthController() controller.AuthController {
	return controller.NewAuthController(r.NewAuthInteractor())
}

func (r *register) NewAuthInteractor() interactor.AuthInteractor {
	return interactor.NewAuthInteractor(r.NewAuthRepository())
}

func (r *register) NewAuthRepository() ur.AuthRepository {
	return ir.NewAuthRepository(r.db)
}
