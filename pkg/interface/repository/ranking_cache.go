package repository

import (
	"context"
	"strings"

	"github.com/OrlandoRomo/go-ambassador/pkg/usercase/repository"
	"github.com/go-redis/redis/v8"
)

type rankingCache struct {
	redis *redis.Client
}

func NewRankingCache(r *redis.Client) repository.RankingCache {
	return &rankingCache{r}
}

func (r *rankingCache) GetRanking() (map[string]float64, error) {
	rankings, err := r.redis.ZRevRangeByScoreWithScores(context.Background(), "rankings", &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	if err != nil {
		return nil, err
	}
	results := make(map[string]float64, 0)
	for _, ranking := range rankings {
		member := strings.ToLower(ranking.Member.(string))
		results[member] = ranking.Score
	}

	return results, nil
}
