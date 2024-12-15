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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&model.Error{Detail: "Not found user"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := h.service.DeleteUserById(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&model.Error{Detail: "Not Found"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) PatchUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var userBody *model.UserPatch
	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Error{Detail: "Invalid body"})
		return
	}

	user, err := h.service.PatchUserById(id, userBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Error{Detail: "Failed patch user"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Error{Detail: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
