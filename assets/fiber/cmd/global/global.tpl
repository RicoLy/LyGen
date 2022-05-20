package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	MysqlDb *gorm.DB
	Redis   *redis.Client
)
