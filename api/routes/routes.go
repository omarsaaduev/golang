package routes

import (
	"github.com/gorilla/mux"
	"golang/api/handlers"
	"golang/internal/service"

	"golang/internal/repository"
	"log"
)

// Настройка маршрутов для API
func SetupRouter() *mux.Router {
	// Подключаемся к базе данных

	dbConn, err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Создаем репозиторий и use case
	userRepo := repository.NewUserRepository(dbConn)
	userUseCase := service.NewUserUseCase(userRepo)

	// Создаем обработчик
	userHandler := handler.NewUserHandler(userUseCase)

	// Создаем роутер
	r := mux.NewRouter()

	// Настроим маршруты
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserByID).Methods("GET")

	return r
}
