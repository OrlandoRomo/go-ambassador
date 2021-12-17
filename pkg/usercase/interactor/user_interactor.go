package interactor

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

type userInteractor struct {
	userRepository repository.UserRepository
}

type UserInteractor interface {
	Create(*model.User) error
	Get(*model.User, string) error
	Update(*model.User) error
}

func NewUserInteractor(r repository.UserRepository) UserInteractor {
	return &userInteractor{r}
}

func (u *userInteractor) Create(us *model.User) error {
	err := u.userRepository.CreateUser(us)
	if err != nil {
		return err
	}
	return err
}

func (u *userInteractor) Get(us *model.User, searchType string) error {
	err := u.userRepository.GetUser(us, searchType)
	if err != nil {
		return err
	}
	return nil
}

func (u *userInteractor) Update(us *model.User) error {
	err := u.userRepository.UpdateUser(us)
	if err != nil {
		return err
	}
	return nil
}

func (u *userInteractor) NewPassword(us *model.User) error {
	err := u.userRepository.UpdateUser(us)
	if err != nil {
		return err
	}
	return nil
}
