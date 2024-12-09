package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserCache struct {
	client *redis.Client
}

func NewUserCache(addr, password string, db int) *UserCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &UserCache{client: client}
}

// GetUser Получение данных из Redis
func (c *UserCache) GetUser(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

// SetUser сохраняет данные в Redis
func (c *UserCache) SetUser(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *UserCache) DeleteUser(ctx context.Context, key string) (int64, error) {
	return c.client.Del(ctx, key).Result()
}
