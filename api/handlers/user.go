package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang/internal/model"
	"golang/internal/service"
	"net/http"
	"strconv"
)

// UserHandler представляет обработчик для пользователей
type UserHandler struct {
	useCase *service.UserUseCase
}

// Конструктор UserHandler
func NewUserHandler(useCase *service.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

// Метод для создания пользователя
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdUser, err := h.useCase.CreateUser(user.FirstName, user.LastName, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// Метод для получения пользователя по ID
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.useCase.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
