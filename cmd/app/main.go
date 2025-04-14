package main

import (
	"fmt"
	"log"
	"tasks-api/internal/core/services"
	"tasks-api/internal/http/config"
	"tasks-api/internal/http/handlers"
	"tasks-api/internal/http/router"
	"tasks-api/internal/infra/postgres"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	cfg := config.LoadConfig()
	connStr := cfg.GetConnString()

	pool, err := postgres.ConnectDatabase(connStr)
	if err != nil {
		log.Fatal("Couldn't connect to database")
	}

	taskRepo := postgres.NewTaskRepository(pool)
	taskServices := services.NewTaskServices(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskServices)
	userRepo := postgres.NewUserRepository(pool)
	userServices := services.NewUserServices(userRepo)
	authServices := services.NewAuthService(userServices)
	authHandler := handlers.NewAuthHandler(authServices)

	e := router.NewRouter(*taskHandler, *authHandler)

	fmt.Print("Starting server")
	startServer(e, cfg.ServerConfig.Port)
}

func startServer(e *echo.Echo, port string) {

	err := e.Start(":" + port)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
