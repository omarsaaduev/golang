package handler

import (
	"encoding/json"
	"golang/internal/model"
	"golang/internal/service"
	"net/http"
)

// UserHandler представляет обработчик для пользователей
type UserHandler struct {
	service *service.UserService
}

// Конструктор UserHandler
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Метод для создания пользователя
type Address struct {
	City  string `json:"city"`
	Phone string `json:"phone"`
}

type UserBody struct {
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email"`
	Address   []Address `json:"address"`
}

// CreateUser
// @Summary
// @Description
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserBody true "Request body"
// @Success 201 {object} model.User "Created"
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdUser, err := h.service.CreateUser(user.FirstName, user.LastName, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
