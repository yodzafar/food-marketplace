package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/yodzafar/food-marketpalce/user-service/config"
)

func newConnection(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
	})

	if err := rdb.Ping().Err(); err != nil {
		return nil, fmt.Errorf("redis connect: %w", err)
	}

	return rdb, nil
}

var ProviderSet = wire.NewSet(newConnection)
