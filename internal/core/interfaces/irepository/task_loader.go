package irepository

import (
	"context"
	"tasks-api/internal/core/domain"
)

type TaskLoader interface {
	InsertTask(ctx context.Context, task domain.Task) (*domain.Task, error)
	FetchTasks(ctx context.Context) ([]domain.Task, error)
	GetTaskByID(ctx context.Context, id int) (*domain.Task, error)
	UpdateTask(ctx context.Context, id int) error
	DeleteTask(ctx context.Context, id int) error
}
