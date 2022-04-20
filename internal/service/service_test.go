package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wagaru/task/internal/model"
	"github.com/wagaru/task/internal/repository/mocks"
)

func TestGetTasks(t *testing.T) {
	mockRepo := new(mocks.Repository)
	tasks := []*model.Task{
		{
			ID:     1,
			Name:   "task1",
			Status: 0,
		},
		{
			ID:     2,
			Name:   "task2",
			Status: 1,
		},
	}
	mockRepo.On("GetTasks").Return(tasks, nil).Once()
	svc := NewService(mockRepo)
	result, err := svc.GetTasks()
	assert.NoError(t, err)
	assert.Equal(t, result, tasks)
	mockRepo.AssertExpectations(t)
}

func TestGetTask(t *testing.T) {
	mockRepo := new(mocks.Repository)
	task := &model.Task{
		ID:     1,
		Name:   "task",
		Status: 0,
	}
	mockRepo.On("GetTask", task.ID).Return(task, nil).Once()
	svc := NewService(mockRepo)
	result, err := svc.GetTask(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, result, task)
	mockRepo.AssertExpectations(t)
}
func TestCreateTask(t *testing.T) {
	mockRepo := new(mocks.Repository)
	task := &model.Task{
		ID:     1,
		Name:   "task",
		Status: 0,
	}
	mockRepo.On("CreateTask", task.Name).Return(task, nil).Once()
	svc := NewService(mockRepo)
	result, err := svc.CreateTask(task.Name)
	assert.NoError(t, err)
	assert.Equal(t, result, task)
	mockRepo.AssertExpectations(t)
}
func TestUpdateTask(t *testing.T) {
	mockRepo := new(mocks.Repository)
	task := &model.Task{
		ID:     1,
		Name:   "task",
		Status: 0,
	}
	mockRepo.On("UpdateTask", task.ID, task.Name, task.Status).Return(task, nil).Once()
	svc := NewService(mockRepo)
	result, err := svc.UpdateTask(task.ID, task.Name, task.Status)
	assert.NoError(t, err)
	assert.Equal(t, result, task)
	mockRepo.AssertExpectations(t)
}
func TestDeleteTask(t *testing.T) {
	id := uint32(1)
	mockRepo := new(mocks.Repository)
	mockRepo.On("DeleteTask", id).Return(nil).Once()
	svc := NewService(mockRepo)
	err := svc.DeleteTask(id)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
