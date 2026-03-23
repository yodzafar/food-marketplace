package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yodzafar/food-marketpalce/user-service/internal/domain"
)

const (
	userCachePrefix = "user:"
	userCacheTTL    = 30 * time.Minute
)

type userCacheRepo struct {
	redis *redis.Client
}

func NewUserCache(redis *redis.Client) domain.UserCache {
	return &userCacheRepo{redis: redis}
}

func (r *userCacheRepo) Set(ctx context.Context, user *domain.User) error {
	data, err := json.Marshal(user)

	if err != nil {
		return fmt.Errorf("failed to marshal user: %w", err)
	}
	key := r.key(user.ID)

	if err := r.redis.Set(ctx, key, data, userCacheTTL).Err(); err != nil {
		return fmt.Errorf("failed to set user in cache: %w", err)
	}

	return nil
}

func (r *userCacheRepo) Get(ctx context.Context, id string) (*domain.User, error) {
	key := r.key(id)

	data, err := r.redis.Get(ctx, key).Bytes()

	if errors.Is(err, redis.Nil) {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get user from cache: %w", err)
	}

	user := &domain.User{}
	if err := json.Unmarshal(data, user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return user, nil
}

func (r *userCacheRepo) Delete(ctx context.Context, id string) error {
	key := r.key(id)

	if err := r.redis.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete user from cache: %w", err)
	}

	return nil
}

func (r *userCacheRepo) key(id string) string {
	return userCachePrefix + id
}
