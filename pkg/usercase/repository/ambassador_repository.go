package repository

import "github.com/OrlandoRomo/go-ambassador/pkg/domain/model"

type AmbassadorRepository interface {
	GetAmbassadors() ([]*model.User, error)
}
