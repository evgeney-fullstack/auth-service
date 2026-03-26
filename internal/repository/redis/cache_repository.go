package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Authorization interface {
	Get(ctx context.Context, key string, dest interface{}) error
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, keys ...string) error
}

type CacheRepository struct {
	Authorization
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{
		Authorization: NewRedisRepository(client),
	}
}
