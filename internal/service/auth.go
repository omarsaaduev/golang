package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang/internal/model/auth"
	"golang/internal/repository"
	"golang/pkg/utils"
	"log/slog"
	"net/smtp"
	"os"
	"time"
)

type AuthService struct {
	repo        *repository.AuthRepository
	verifyCodes map[string]string // Временное хранилище для email -> verification code
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{
		repo:        repo,
		verifyCodes: make(map[string]string), // Инициализация
	}
}

func (s *AuthService) generateVerificationCode() string {
	code := fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
	slog.Info("EXPCODE", code)
	return code // 6-значный код
}

func (s *AuthService) sendCode(email string) error {
	code := s.generateVerificationCode()

	// Сохраняем код во временное хранилище
	s.verifyCodes[email] = code

	// Настройка и отправка письма
	from := os.Getenv("MAIL_FROM")
	password := os.Getenv("MAIL_PASSWORD")
	to := []string{email}

	// SMTP server details
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Message content
	subject := "Subject: Test Email from Go\n"
	body := fmt.Sprintf("Your verificatyon code: %s.\n", code)
	message := []byte(subject + "\n" + body)

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		slog.Error("Failed to send email: %v", err)
		return err
	}

	slog.Info("Email sent successfully!")
	return nil
}

func (s *AuthService) SignUp(user *auth.UserCreate) (*auth.TokenDetails, error) {
	// Проверяем, существует ли пользователь
	exists, _ := s.repo.CheckUserExists(user.Email)
	if exists {
		return nil, errors.New("user already exists")
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Сохраняем пользователя в базе
	userID, err := s.repo.CreateUser(user.Email, user.LastName, user.Email, string(hashedPassword), user.CreatedAt)
	if err != nil {
		return nil, err
	}

	// Генерируем токены
	authUtils := utils.NewAuthUtils()
	tokens, err := authUtils.CreateToken(userID)
	if err != nil {
		return nil, err
	}

	// Сохраняем токены в базе
	err = s.repo.SaveTokenData(tokens.AtExpires, tokens.RtExpires, userID)
	if err != nil {
		return nil, err
	}

	// Отправляем верификационный код
	err = s.sendCode(user.Email)
	if err != nil {
		slog.Info("ОТПРАВКА КОДА:", err)
	}

	return tokens, nil
}

func (s *AuthService) VerifyCode(email, code string) error {
	// Получаем код из временного хранилища
	expectedCode, exists := s.verifyCodes[email]
	if !exists {
		return errors.New("verification code not found for the given email")
	}

	// Сравниваем коды
	if expectedCode != code {
		return errors.New("verification code does not match")
	}

	// Активируем пользователя в базе данных
	err := s.repo.ActivateUser(email)
	if err != nil {
		return err
	}

	// Удаляем код из хранилища
	delete(s.verifyCodes, email)
	return nil
}
