package cache

import (
	"github.com/compscore/compscore/pkg/config"
	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
)

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Url,
		Password: config.Redis.Password,
		DB:       0,
	})
}
