package interactor

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
)

type linkInteractor struct {
	linkRepository repository.LinkRepository
}

type LinkInteractor interface {
	Create(*model.Link) error
}

func NewLinkInteractor(l repository.LinkRepository) LinkInteractor {
	return &linkInteractor{l}
}

func (l *linkInteractor) Create(link *model.Link) error {
	err := l.linkRepository.CreateLink(link)
	if err != nil {
		return err
	}
	return nil
}
