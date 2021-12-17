package interactor

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

type rankingInteractor struct {
	rankingCache repository.RankingCache
}

type RankingInteractor interface {
	Get() (map[string]float64, error)
}

func NewRankingInteractor(c repository.RankingCache) RankingInteractor {
	return &rankingInteractor{c}
}

func (a *rankingInteractor) Get() (map[string]float64, error) {
	ambassadors, err := a.rankingCache.GetRanking()
	if err != nil {
		return nil, err
	}
	return ambassadors, nil
}
