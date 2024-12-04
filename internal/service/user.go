package service

import (
	"golang/internal/model"
	"golang/internal/repository"
	"time"
)

// Интерфейс для usecase
type UserService struct {
	repo *repository.UserRepository
}

// Конструктор UserUseCase
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Метод для создания пользователя
func (us *UserService) CreateUser(user *model.User) (*model.User, error) {
	user.CreatedAt = time.Now()
	err := us.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Метод для получения пользователя
func (us *UserService) GetUserByID(id string) (*model.User, error) {
	user, err := us.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
