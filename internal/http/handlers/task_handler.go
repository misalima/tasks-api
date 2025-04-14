package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"tasks-api/internal/core/domain"
	"tasks-api/internal/core/interfaces/iservice"
	dto2 "tasks-api/internal/http/dto"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service iservice.TaskManager
}

func NewTaskHandler(service iservice.TaskManager) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var createTask dto2.CreateTaskDTO
	if err := c.Bind(&createTask); err != nil {
		log.Println("Error parsing request: ", err)
		return c.JSON(http.StatusBadRequest, dto2.ErrorResponse{
			Message: "Invalid request data",
		})
	}

	if createTask.Title == "" || createTask.Description == "" {
		return c.JSON(http.StatusBadRequest, dto2.ErrorResponse{
			Message: "Invalid request data",
		})
	}

	task, err := h.service.CreateTask(c.Request().Context(), createTask.ToDomain())
	if err != nil {
		log.Println("Error while creating task: ", err)
		return c.JSON(http.StatusInternalServerError, dto2.ErrorResponse{Message: "Internal Server Error"})
	}

	response := dto2.FromDomain(*task)

	return c.JSON(http.StatusCreated, response)
}

func (h *TaskHandler) GetTaskByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto2.ErrorResponse{Message: "Invalid request data"})
	}
	task, err := h.service.GetTaskByID(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
			return c.JSON(http.StatusNotFound, dto2.ErrorResponse{Message: "Task not found"})
		}
		return c.JSON(http.StatusInternalServerError, dto2.ErrorResponse{Message: "Internal Server Error"})
	}
	response := dto2.FromDomain(*task)
	return c.JSON(http.StatusOK, response)
}
