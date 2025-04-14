package iservice

import (
	"context"
	"tasks-api/internal/core/domain"
)

type TaskManager interface {
	CreateTask(ctx context.Context, task domain.Task) (*domain.Task, error)
	ListTasks(ctx context.Context) ([]domain.Task, error)
	GetTaskByID(ctx context.Context, id int) (*domain.Task, error)
	UpdateTask(ctx context.Context, id int) error
	DeleteTask(ctx context.Context, id int) error
	SetTaskDone(ctx context.Context, id int) error
}
