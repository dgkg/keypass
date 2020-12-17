package cache

import "context"

type CacheDB interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) ([]byte, error)
}
