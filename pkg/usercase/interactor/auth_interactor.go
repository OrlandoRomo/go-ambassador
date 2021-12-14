package interactor

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

type AuthInteractor interface {
	Login(email string) (*model.User, error)
}

type authInteractor struct {
	authRepository repository.AuthRepository
}

func NewAuthInteractor(r repository.AuthRepository) AuthInteractor {
	return &authInteractor{r}
}

func (a *authInteractor) Login(email string) (*model.User, error) {
	user, err := a.authRepository.Login(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
