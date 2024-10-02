package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type IStorageStrategy interface {
	Set(ctx context.Context, key string, data any) error
	Get(ctx context.Context, key string) (string, error)
}

// Implements IStorageStrategy
type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(addr string) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisStorage{client: client}
}

func (r *RedisStorage) Set(ctx context.Context, key string, data any) error {
	return r.client.Set(ctx, key, data, 0).Err()
}

func (r *RedisStorage) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
