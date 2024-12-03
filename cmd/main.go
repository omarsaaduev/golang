// @title My API
// @version 1.0
// @description This is a sample API documentation using Swagger.
// @host localhost:8080
// @BasePath /
package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang/api/routes"
	_ "golang/docs"
	"golang/pkg/logger"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	errEnv := godotenv.Load("../.env")
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	logg := logger.SetupLogger()
	if logg == nil {
		fmt.Println("Logger is nil")
	}
	slog.SetDefault(logg)

	r := routes.SetupRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		slog.Error("Error starting server")
		return
	}
	slog.Error("Server starting...")

}
