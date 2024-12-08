package handlers

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

// NewUserHandler Конструктор
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetUserByID Метод для получения пользователя по id
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.service.GetUser(r.Context(), id)
	slog.Info("user", user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&model.Error{Detail: "Not found user"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
