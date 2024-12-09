package redis

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

type Config struct {
	Host string
	Port uint
}

func NewRedisClient(cfg Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return client, nil
}
