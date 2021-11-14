package controller

import (
	"context"
	"net/http"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func GetRankings(c *fiber.Ctx) error {
	rankings, err := database.Cache.ZRevRangeByScoreWithScores(context.Background(), "rankings", &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	results := make(map[string]float64)

	for _, ranking := range rankings {
		results[ranking.Member.(string)] = ranking.Score
	}

	return c.JSON(fiber.Map{
		"rankings": rankings,
	})
}
