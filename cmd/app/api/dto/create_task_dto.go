package dto

import "tasks-api/internal/core/domain"

type CreateTaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (dto *CreateTaskDTO) ToDomain() domain.Task {
	return domain.Task{
		Title: dto.Title,
		Description: dto.Description,
		Status: dto.Status,
	}
}

