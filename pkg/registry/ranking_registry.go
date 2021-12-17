package registry

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	ir "github.com/OrlandoRomo/go-ambassador/pkg/interface/repository"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	ur "github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

func (r *register) NewRankingController() controller.RankingController {
	return controller.NewRankingController(r.NewRankingInteractor())
}

func (r *register) NewRankingInteractor() interactor.RankingInteractor {
	return interactor.NewRankingInteractor(r.NewRankingCache())
}

func (r *register) NewRankingCache() ur.RankingCache {
	return ir.NewRankingCache(r.redis)

}
