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
func (u *UserCache) GetUser(ctx context.Context, key string) (string, error) {
	return u.client.Get(ctx, key).Result()
}

// SetUser сохраняет данные в Redis
func (u *UserCache) SetUser(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return u.client.Set(ctx, key, value, ttl).Err()
}
