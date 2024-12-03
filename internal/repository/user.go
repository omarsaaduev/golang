package repository

import (
	"database/sql"
	"fmt"
	"golang/internal/model"
)

// Репозиторий для User
type UserRepository struct {
	db *sql.DB
}

// Конструктор UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Создание нового пользователя
func (r *UserRepository) Create(user *model.User) error {
	query := `INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}
