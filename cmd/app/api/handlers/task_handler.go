package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"tasks-api/cmd/app/api/dto"
	"tasks-api/internal/core/domain"
	"tasks-api/internal/core/interfaces/iservice"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service iservice.TaskManager
}

func NewTaskHandler(service iservice.TaskManager) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var createTask dto.CreateTaskDTO
	if err := c.Bind(&createTask); err != nil {
		log.Println("Error parsing request: ", err)
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request data",
		})
	}

	if createTask.Title == "" || createTask.Description == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request data",
		})
	}

	task, err := h.service.CreateTask(c.Request().Context(), createTask.ToDomain())
	if err != nil {
		log.Println("Error while creating task: ", err)
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal Server Error"})
	}
	 
	response := dto.FromDomain(*task)
	
	return c.JSON(http.StatusCreated, response)
}

func (h *TaskHandler) GetTaskByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request data"})
	}
	task, err := h.service.GetTaskByID(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Task not found"})
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal Server Error"})
	}
	response := dto.FromDomain(*task)
	return c.JSON(http.StatusOK, response)
}