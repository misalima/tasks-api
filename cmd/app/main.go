package main

import (
	"fmt"
	"log"
	"tasks-api/cmd/app/api/config"
	"tasks-api/cmd/app/api/handlers"
	"tasks-api/cmd/app/api/router"
	"tasks-api/internal/core/services"
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
	
	
	pool, err := postgres.ConnectDatabase(connStr);
	if err != nil {
		log.Fatal("Couldn't connect to database")
	}

	taskRepo := postgres.NewTaskRepository(pool)
	taskServices := services.NewTaskServices(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskServices)
	
	e := router.NewRouter(*taskHandler)

	fmt.Print("Starting server")
	startServer(e, cfg.ServerConfig.Port)
}

func startServer(e *echo.Echo, port string) {
	
	err := e.Start(":" + port)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
