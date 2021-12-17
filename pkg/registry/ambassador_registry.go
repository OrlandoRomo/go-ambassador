package registry

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	ir "github.com/OrlandoRomo/go-ambassador/pkg/interface/repository"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	ur "github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

func (r *register) NewAmbassadorController() controller.AmbassadorController {
	return controller.NewAmbassadorController(r.NewAmbassadorInteractor())
}

func (r *register) NewAmbassadorInteractor() interactor.AmbassadorInteractor {
	return interactor.NewAmbassadorInteractor(r.NewAmbassadorRepository())
}

func (r *register) NewAmbassadorRepository() ur.AmbassadorRepository {
	return ir.NewAmbassadorRepository(r.db)
}
