package registry

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	ir "github.com/OrlandoRomo/go-ambassador/pkg/interface/repository"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	ur "github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

func (r *register) NewLinkController() controller.LinkController {
	return controller.NewLinkController(r.NewLinkInteractor())
}

func (r *register) NewLinkInteractor() interactor.LinkInteractor {
	return interactor.NewLinkInteractor(r.NewLinkRepository())
}

func (r *register) NewLinkRepository() ur.LinkRepository {
	return ir.NewLinkRepository(r.db)
}
