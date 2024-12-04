package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang/internal/model"
	"golang/internal/service"
	"log/slog"
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

// CreateUser
// @Summary
// @Description
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "Request body"
// @Success 201 {object} model.User "Created"
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Error{
			Detail: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// Метод для получения пользователя по id
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.service.GetUserByID(id)
	slog.Info("user", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
