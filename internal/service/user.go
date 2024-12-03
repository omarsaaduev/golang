package service

import (
	"golang/internal/model"
	"golang/internal/repository"
)

// Интерфейс для usecase
type UserUseCase struct {
	repo *repository.UserRepository
}

// Конструктор UserUseCase
func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

// Метод для создания пользователя
func (uc *UserUseCase) CreateUser(first_name, last_name, email string) (*model.User, error) {
	user := &model.User{FirstName: first_name, LastName: last_name, Email: email}
	err := uc.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Метод для получения пользователя по ID
func (uc *UserUseCase) GetUserByID(id int64) (*model.User, error) {
	return uc.repo.GetByID(id)
}
