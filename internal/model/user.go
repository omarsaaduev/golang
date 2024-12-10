package model

import "time"

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserPatch struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
