package repository

import (
	"database/sql"
	"errors"
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

// GetUserById Получение пользвоателя по id
func (r *UserRepository) GetUserById(id string) (*model.User, error) {
	var user *model.User
	user = &model.User{}
	query := `SELECT id, first_name, last_name, email, created_at FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)

	if err != nil {

		slog.Error("Error getting user by id:", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUserById(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
