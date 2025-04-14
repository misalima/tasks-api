package services

import (
	"context"
	"tasks-api/internal/core/domain"
	"tasks-api/internal/core/interfaces/irepository"
	"tasks-api/internal/core/interfaces/iservice"
)

var _ iservice.TaskManager = (*TaskServices)(nil)

type TaskServices struct {
	repo irepository.TaskLoader
}

func NewTaskServices(repo irepository.TaskLoader) *TaskServices {
	return &TaskServices{repo: repo}
}

func (s *TaskServices) CreateTask(ctx context.Context, task domain.Task) (*domain.Task, error) {
	createdTask, err := s.repo.InsertTask(ctx, task)
	if err != nil {
		return nil, err
	}

	createdTask.Status = "pendente"

	return createdTask, nil
}

func (s *TaskServices) ListTasks(ctx context.Context) ([]domain.Task, error) {
	panic("Not implemented")
}
func (s *TaskServices) GetTaskByID(ctx context.Context, id int) (*domain.Task, error) {
	return s.repo.GetTaskByID(ctx, id)
}
func (s *TaskServices) UpdateTask(ctx context.Context, id int) error {
	panic("Not implemented")
}
func (s *TaskServices) DeleteTask(ctx context.Context, id int) error {
	panic("Not implemented")
}
func (s *TaskServices) SetTaskDone(ctx context.Context, id int) error {
	panic("Not implemented")
}
