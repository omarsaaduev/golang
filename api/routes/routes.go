package routes

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang/api/handlers"

	"golang/internal/repository"
	"golang/internal/service"
	"log"
)

// Настройка маршрутов для API
func SetupRouter() *mux.Router {
	// Подключаемся к базе данных
	dbConn, err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Создаем репозиторий и сервис юзера
	userRepo := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Создаем репозиторий и сервис авторизации
	authRepo := repository.NewAuthRepository(dbConn)
	authService := service.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Создаем роутер
	r := mux.NewRouter()

	// Настроим маршруты

	//Auth
	r.HandleFunc("/auth/sign-up", authHandler.CreateUser).Methods("POST")
	r.HandleFunc("/auth/confirm", authHandler.VerificationCode).Methods("POST")

	//Users
	// r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
