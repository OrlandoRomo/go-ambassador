package interactor

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

type ambassadorInteractor struct {
	ambassadorRepository repository.AmbassadorRepository
}

type AmbassadorInteractor interface {
	Get() ([]*model.User, error)
}

func NewAmbassadorInteractor(r repository.AmbassadorRepository) AmbassadorInteractor {
	return &ambassadorInteractor{r}
}

func (a *ambassadorInteractor) Get() ([]*model.User, error) {
	ambassadors, err := a.ambassadorRepository.GetAmbassadors()
	if err != nil {
		return nil, err
	}
	return ambassadors, nil
}
