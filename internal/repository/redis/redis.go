package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string
	Port     string
	Password string
	DB       int
}

func NewRedisDB(cfg Config) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
