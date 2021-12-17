package cache

import "github.com/go-redis/redis/v8"

func NewCache(port string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "redis:" + port,
		DB:   0,
	})
}
