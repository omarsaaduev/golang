package service

import (
	"context"
	"encoding/json"
	"golang/internal/cache"
	"golang/internal/model"
	"golang/internal/repository"
	"time"
)

type UserService struct {
	repo  *repository.UserRepository
	cache *cache.UserCache
}

func NewUserService(repo *repository.UserRepository, cache *cache.UserCache) *UserService {
	return &UserService{repo: repo, cache: cache}
}

// GetUser получает пользователя из Redis или PostgresSQL
func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	cacheKey := "user:" + id
	// Получаем юзера из Redis если он есть
	cachedUser, err := s.cache.GetUser(ctx, cacheKey)
	if err == nil {
		var user model.User
		if jsonErr := json.Unmarshal([]byte(cachedUser), &user); jsonErr == nil {
			return &user, nil
		}
	}

	// Если в кэше нет, берём из базы
	user, err := s.repo.GetUserById(id)
	if err != nil || user == nil {
		return nil, err
	}

	// Кэшируем данные
	userData, _ := json.Marshal(user)
	_ = s.cache.SetUser(ctx, cacheKey, string(userData), 60*time.Second)
	return user, nil
}
