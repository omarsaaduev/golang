package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang/internal/service"
	"log/slog"
	"net/http"
)

// UserHandler представляет обработчик для пользователей
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler Конструктор
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetUserByID Метод для получения пользователя по id
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
