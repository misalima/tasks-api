package dto

import "tasks-api/internal/core/domain"

type TaskResponseDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func FromDomain(task domain.Task) TaskResponseDTO {
	return TaskResponseDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
}
