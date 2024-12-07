package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

// CheckUserExists проверяет, существует ли пользователь с указанным email
func (r *AuthRepository) CheckUserExists(email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	return exists, err
}

// CreateUser создает нового пользователя
func (r *AuthRepository) CreateUser(firstName, lastName, email, password string, createdAt time.Time) (string, error) {
	// Используем RETURNING для получения ID
	var id string
	query := `
		INSERT INTO users (first_name, last_name, created_at, email, password) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id`

	// Выполняем запрос и получаем значение id
	err := r.db.QueryRow(query, firstName, lastName, createdAt, email, password).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

// SaveTokenData сохраняет срок действия токенов в базу данных
func (r *AuthRepository) SaveTokenData(atExpires, rtExpires int64, userID string) error {
	query := `
		UPDATE users 
		    SET access_exp = $1, refresh_exp = $2 
		WHERE id = $3`
	_, err := r.db.Exec(query, atExpires, rtExpires, userID)
	return err
}

func (r *AuthRepository) ActivateUser(email string) error {
	query := `UPDATE users SET is_active = true WHERE email = $1`
	_, err := r.db.Exec(query, email)
	if err != nil {
		return fmt.Errorf("failed to activate user: %v", err)
	}
	return nil
}
