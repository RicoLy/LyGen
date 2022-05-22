package global

import (
	"{{.ProjectName}}/cmd/internal/config"
)


func InitRedis(c *config.Config) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}