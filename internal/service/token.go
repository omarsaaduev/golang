package service

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang/internal/model/auth"
	"time"
)

// Секреты для подписи токенов
var accessSecret = []byte("your_access_secret")   // Секрет для Access токена
var refreshSecret = []byte("your_refresh_secret") // Секрет для Refresh токена

type TokenService struct{}

// NewTokenService Конструктор TokenService
func NewTokenService() *TokenService {
	return &TokenService{}
}

// CreateToken генерирует пару Access и Refresh токенов
func (t *TokenService) CreateToken(userID string) (*auth.TokenDetails, error) {
	td := &auth.TokenDetails{}

	// Устанавливаем срок действия Access токена (15 минут)
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.New().String() // Уникальный идентификатор для Access токена

	// Устанавливаем срок действия Refresh токена (7 дней)
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.New().String() // Уникальный идентификатор для Refresh токена

	// Генерация Access токена
	atClaims := jwt.MapClaims{
		"authorized":  true,
		"access_uuid": td.AccessUUID,
		"user_id":     userID,
		"exp":         td.AtExpires, // Время истечения токена
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	var err error
	td.AccessToken, err = at.SignedString(accessSecret) // Подписываем токен секретом
	if err != nil {
		return nil, err
	}

	// Генерация Refresh токена
	rtClaims := jwt.MapClaims{
		"refresh_uuid": td.RefreshUUID,
		"user_id":      userID,
		"exp":          td.RtExpires,
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString(refreshSecret) // Подписываем токен секретом
	if err != nil {
		return nil, err
	}

	return td, nil
}
