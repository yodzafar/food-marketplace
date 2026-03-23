package redis

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/yodzafar/food-marketpalce/user-service/config"
)

func New(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("redis connect: %w", err)
	}

	return rdb, nil
}

var ProviderSet = wire.NewSet(New)
