package cache

import (
	"context"
	"time"
)

type RedisClient interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Close() error
}
