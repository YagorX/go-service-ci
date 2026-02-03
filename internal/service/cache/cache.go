package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cacheable interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}

type Cache[T Cacheable] struct {
	r *redis.Client
}

func New[T Cacheable](rdb *redis.Client) *Cache[T] {
	return &Cache[T]{r: rdb}
}

func (c *Cache[T]) Set(ctx context.Context, key string, value T, expiration time.Duration) error {
	encodedValue, err := value.Encode()
	if err != nil {
		return err
	}

	err = c.r.Set(ctx, key, encodedValue, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache[T]) SetMany(ctx context.Context, items map[string]T, expiration time.Duration) error {
	pipe := c.r.TxPipeline()

	for key, value := range items {
		encodedValue, err := value.Encode()
		if err != nil {
			return err
		}

		pipe.Set(ctx, key, encodedValue, expiration)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache[T]) Get(ctx context.Context, key string, destination T) error {
	resBytes, err := c.r.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	err = destination.Decode(resBytes)
	if err != nil {
		return err
	}

	return nil
}
