package handlers

import (
	"encoding/json"
	"fmt"
	"golang/internal/model"
	"golang/internal/model/auth"
	"golang/internal/service"
	"net/http"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *auth.UserCreate
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(&model.Error{
			Detail: "Invalid Data",
		})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	createdUser, err := h.service.SignUp(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Error{
			Detail: fmt.Sprintf("%v", err),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&auth.TokenResponse{
		AccessToken:  createdUser.AccessToken,
		RefreshToken: createdUser.RefreshToken,
	})
}

func (h *AuthHandler) VerificationCode(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request data"})
		return
	}

	// Проверяем код
	err := h.service.VerifyCode(request.Email, request.Code)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Успешная верификация
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "verification successful"})
}
