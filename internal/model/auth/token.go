package auth

// TokenDetails содержит информацию о сгенерированных токенах
type TokenDetails struct {
	AccessToken  string // Access токен
	RefreshToken string // Refresh токен
	AccessUUID   string // Уникальный идентификатор для Access токена
	RefreshUUID  string // Уникальный идентификатор для Refresh токена
	AtExpires    int64  // Время жизни Access токена (в секундах)
	RtExpires    int64  // Время жизни Refresh токена (в секундах)
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type VerificationResponse struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
