package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	name := os.Getenv("POSTGRES_NAME")
	host := os.Getenv("POSTGRES_HOST")

	slog.Info(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name))
	dbStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	db, err := sql.Open("postgres", dbStr)
	if err != nil {
		errMsg := fmt.Sprintf("Error connecting to database: %v", err)
		slog.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	// Проверка подключения
	if err := db.Ping(); err != nil {
		errMsg := fmt.Sprintf("Error pinging database: %v", err)
		slog.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	return db, nil
}
