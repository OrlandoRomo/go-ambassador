package registry

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	ir "github.com/OrlandoRomo/go-ambassador/pkg/interface/repository"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	ur "github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

func (r *register) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *register) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository())
}

func (r *register) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}
