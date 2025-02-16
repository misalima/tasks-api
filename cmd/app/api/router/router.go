package router

import (
	"net/http"
	"tasks-api/cmd/app/api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(taskHandler handlers.TaskHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodOptions,
			http.MethodDelete,
		},
	}))
	
	api := e.Group("/api")

	api.POST("/tasks", taskHandler.CreateTask)
	api.GET("/tasks/:id", taskHandler.GetTaskByID)

	return e
}