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
	query := `INSERT INTO users (first_name, last_name, email) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

//Получение пользователя по ID

func (r *UserRepository) GetByID(id int64) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, first_name, last_name, email FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("could not get user: %w", err)
	}
	return user, nil
}
