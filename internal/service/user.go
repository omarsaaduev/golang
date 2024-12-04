package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang/internal/model"
	"golang/internal/repository"
	"log/slog"
	"time"
)

// UserService Интерфейс для userService
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService Конструктор UserUseCase
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser Метод для создания пользователя
func (us *UserService) CreateUser(user *model.User) (*model.User, error) {
	user.CreatedAt = time.Now()

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		slog.Error("Email validation error")
		return nil, fmt.Errorf("email validation error")
	}

	err = us.repo.Create(user)
	if err != nil {
		slog.Error("Failed to create user")
		return nil, fmt.Errorf("failed to create user")
	}

	return user, nil
}

// GetUserByID Метод для получения пользователя
func (us *UserService) GetUserByID(id string) (*model.User, error) {
	user, err := us.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
