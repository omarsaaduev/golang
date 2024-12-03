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

	// Создаем репозиторий и сервис
	userRepo := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo)

	// Создаем обработчик
	userHandler := handler.NewUserHandler(userService)

	// Создаем роутер
	r := mux.NewRouter()

	// Настроим маршруты
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
