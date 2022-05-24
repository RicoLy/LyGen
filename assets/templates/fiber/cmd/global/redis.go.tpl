package global

import (
	"{{.}}/cmd/internal/config"
	"github.com/go-redis/redis/v8"
)


func InitRedis(c *config.Config) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}