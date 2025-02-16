package services

import (
	"context"
	"errors"
	"tasks-api/internal/core/domain"
	mock_irepository "tasks-api/internal/core/interfaces/irepository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



func TestGetTaskByID_Success(t *testing.T) {
	mockRepo := new(mock_irepository.MockTaskRepository)
	service := TaskServices{repo: mockRepo}

	mockRepo.On("GetTaskByID", mock.Anything, 1).Return(&domain.Task{
		ID: 1,
		Title: "Do homework",
		Description: "Geography and History homework",
		Status: "pendente",
	}, nil)

	task, err := service.GetTaskByID(context.Background(), 1)

	assert.Nil(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, task.ID, 1)
	assert.Equal(t, task.Title, "Do homework")

	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_NotFound(t *testing.T) {
	mockRepo := new(mock_irepository.MockTaskRepository)
	service := TaskServices{repo: mockRepo}

	mockRepo.On("GetTaskByID", mock.Anything, 1).Return(nil, errors.New("task not found"))

	task, err := service.GetTaskByID(context.Background(), 1)

	assert.NotNil(t, err)
	assert.Nil(t, task)
	assert.Equal(t, "task not found", err.Error())

	mockRepo.AssertExpectations(t,)
}