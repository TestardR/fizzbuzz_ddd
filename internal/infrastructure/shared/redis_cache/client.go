package redis_cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	maxKey  = "max_key"
	maxHits = "max_hits"
)

type Client struct {
	redis *redis.Client
}

func NewClient(redis *redis.Client) *Client {
	return &Client{redis: redis}
}

func (c *Client) Set(ctx context.Context, key string, data []byte, expiration time.Duration) error {
	result := c.redis.Set(ctx, key, data, expiration)
	return result.Err()
}

func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	result := c.redis.Get(ctx, key)
	resultBytes, err := result.Bytes()

	return resultBytes, err
}

func (c *Client) Delete(ctx context.Context, key string) error {
	result := c.redis.Del(ctx, key)
	return result.Err()
}

func (c *Client) IncrementCount(ctx context.Context, key string) error {
	result := c.redis.Incr(ctx, key)
	return result.Err()
}

func (c *Client) GetMaxEntries(ctx context.Context) (string, int, error) {
	key, err := c.redis.Get(ctx, maxKey).Result()
	if err != nil {
		return "", 0, err
	}

	hits, err := c.redis.Get(ctx, maxHits).Int()
	if err != nil {
		return "", 0, err
	}

	return key, hits, nil
}

func (c *Client) setMaxEntries(ctx context.Context, keyValue string, hitsValue int) error {
	pipe := c.redis.TxPipeline()
	pipe.Set(ctx, maxKey, keyValue, 0)
	pipe.Set(ctx, maxHits, hitsValue, 0)

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}
