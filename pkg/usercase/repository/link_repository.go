package repository

import "github.com/OrlandoRomo/go-ambassador/pkg/domain/model"

type LinkRepository interface {
	CreateLink() (model.Link, error)
}
