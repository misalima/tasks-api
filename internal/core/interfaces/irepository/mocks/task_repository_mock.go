package mock_irepository

import (
	"context"
	"tasks-api/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetTaskByID(ctx context.Context, id int) (*domain.Task, error) {
	args := m.Called(ctx, id)
	if task, ok := args.Get(0).(*domain.Task); ok {
		return task, args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockTaskRepository) InsertTask(ctx context.Context, task domain.Task) (*domain.Task, error) {
	panic("Not implemented")
}  
func (m *MockTaskRepository) FetchTasks(ctx context.Context, ) ([]domain.Task, error) {
	panic("Not implemented")
}
func (m *MockTaskRepository) UpdateTask(ctx context.Context, id int) error {
	panic("Not implemented")
}
func (m *MockTaskRepository) DeleteTask(ctx context.Context, id int) error {
	panic("Not implemented")
}


	
	