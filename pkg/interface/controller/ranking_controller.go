package controller

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/interactor"
	"github.com/gofiber/fiber/v2"
)

type rankingController struct {
	rankingInteractor interactor.RankingInteractor
}

type RankingController interface {
	GetRankings(c *fiber.Ctx) error
}

func NewRankingController(i interactor.RankingInteractor) RankingController {
	return &rankingController{i}
}

func (r *rankingController) GetRankings(c *fiber.Ctx) error {
	rankings, err := r.rankingInteractor.Get()
	if err != nil {
		return model.EncodeError(c, err)
	}

	return c.JSON(fiber.Map{
		"rankings": rankings,
	})
}
