package repository

import (
	"database/sql"
	"fmt"
	"golang/internal/model"
	"log/slog"
)

// UserRepository Репозиторий для User
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository Конструктор UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create Создание нового пользователя
func (r *UserRepository) Create(user *model.User) error {
	query := `INSERT INTO users (first_name, last_name, email, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.CreatedAt).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

// GetUserById Получение пользвоателя по id
func (r *UserRepository) GetUserById(id string) (*model.User, error) {
	var user *model.User
	user = &model.User{}
	query := `SELECT id, first_name, last_name, email, created_at FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	if err != nil {
		slog.Error("Error getting user by id:", err)
	}
	return user, nil
}
