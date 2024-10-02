package database

import (
	"context"
)

type StorageContext struct {
	strategy IStorageStrategy
}

func NewStorageContext(strategy IStorageStrategy) *StorageContext {
	return &StorageContext{strategy: strategy}
}

func (c *StorageContext) Set(ctx context.Context, key string, value interface{}) error {
	return c.strategy.Set(ctx, key, value)
}

func (c *StorageContext) Get(ctx context.Context, key string) (string, error) {
	return c.strategy.Get(ctx, key)
}
