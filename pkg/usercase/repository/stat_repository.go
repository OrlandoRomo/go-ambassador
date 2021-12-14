package repository

import "github.com/OrlandoRomo/go-ambassador/pkg/domain/model"

type StatRepository interface {
	GetStats() ([]model.Stat, error)
}
