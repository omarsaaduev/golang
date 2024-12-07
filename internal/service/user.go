package service

import (
	"golang/internal/model"
	"golang/internal/repository"
)

// UserService Интерфейс для userService
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService Конструктор UserUseCase
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUserByID Метод для получения пользователя
func (us *UserService) GetUserByID(id string) (*model.User, error) {
	user, err := us.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
