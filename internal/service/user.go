package service

import (
	"golang/internal/model"
	"golang/internal/repository"
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
func (us *UserService) CreateUser(firstName, lastName, email string) (*model.User, error) {
	user := &model.User{FirstName: firstName, LastName: lastName, Email: email}
	err := us.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
