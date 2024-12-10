package routes

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang/api/handlers"
	"golang/internal/cache"
	"golang/internal/repository"
	"golang/internal/service"
	"log"
	"os"
)

// SetupRouter Настройка маршрутов для API
func SetupRouter() *mux.Router {
	// Подключаемся к базе данных
	dbConn, err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Подключение к Redis
	redisCache := cache.NewUserCache(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"), 0)

	// Создаем репозиторий и сервис юзера
	userRepo := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo, redisCache)
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
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.DeleteUserById).Methods("DELETE")
	r.HandleFunc("/users/{id}", userHandler.PatchUserById).Methods("PATCH")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
