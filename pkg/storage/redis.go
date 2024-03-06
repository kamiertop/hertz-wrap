package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis() error {
	Redis = redis.NewClient(&redis.Options{})
	cmd := Redis.Ping(context.Background())

	return cmd.Err()
}
