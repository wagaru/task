package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wagaru/task/config"
	"github.com/wagaru/task/internal/errcode"
	"github.com/wagaru/task/internal/model"
	"github.com/wagaru/task/internal/service/mocks"
)

func TestGetTasks(t *testing.T) {
	mockService := new(mocks.Service)
	t.Run("UnknownError", func(t *testing.T) {
		mockService.On("GetTasks").Return(nil, errors.New("fake")).Once()
		delivery := NewDelivery(mockService, &config.ServerConfig{})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/tasks", nil)
		delivery.engine.ServeHTTP(w, req)
		expected, _ := json.Marshal(map[string]interface{}{
			"code":    errcode.UnknownError.Code(),
			"message": errcode.UnknownError.Message(),
		})
		assert.Equal(t, errcode.UnknownError.StatusCode(), w.Code)
		assert.Equal(t, w.Body.String(), string(expected))
		mockService.AssertExpectations(t)
	})
	t.Run("Success", func(t *testing.T) {
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
		mockService.On("GetTasks").Return(tasks, nil).Once()
		delivery := NewDelivery(mockService, &config.ServerConfig{})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/tasks", nil)
		delivery.engine.ServeHTTP(w, req)
		expected, _ := json.Marshal(map[string]interface{}{
			"result": tasks,
		})
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, w.Body.String(), string(expected))
		mockService.AssertExpectations(t)
	})
}

func TestCreateTask(t *testing.T) {
	mockService := new(mocks.Service)
	t.Run("InvalidParams", func(t *testing.T) {
		delivery := NewDelivery(mockService, &config.ServerConfig{})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/tasks", nil)
		delivery.engine.ServeHTTP(w, req)
		expected, _ := json.Marshal(map[string]interface{}{
			"code":    errcode.InvalidParams.Code(),
			"message": errcode.InvalidParams.Message(),
		})
		assert.Equal(t, errcode.InvalidParams.StatusCode(), w.Code)
		assert.JSONEq(t, w.Body.String(), string(expected))
		mockService.AssertNotCalled(t, "CreateTasks")
	})
	t.Run("DuplicateRecords", func(t *testing.T) {
		taskName := "task"
		mockService.On("CreateTask", taskName).Return(nil, errcode.DuplicateRecords).Once()
		delivery := NewDelivery(mockService, &config.ServerConfig{})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBufferString(fmt.Sprintf(`{"name":"%s"}`, taskName)))
		delivery.engine.ServeHTTP(w, req)
		expected, _ := json.Marshal(map[string]interface{}{
			"code":    errcode.DuplicateRecords.Code(),
			"message": errcode.DuplicateRecords.Message(),
		})
		assert.Equal(t, errcode.DuplicateRecords.StatusCode(), w.Code)
		assert.JSONEq(t, w.Body.String(), string(expected))
		mockService.AssertNotCalled(t, "CreateTasks")
	})
	t.Run("Success", func(t *testing.T) {
		task := &model.Task{
			ID:     1,
			Name:   "task",
			Status: 0,
		}
		mockService.On("CreateTask", task.Name).Return(task, nil).Once()
		delivery := NewDelivery(mockService, &config.ServerConfig{})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBufferString(fmt.Sprintf(`{"name":"%s"}`, task.Name)))
		delivery.engine.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)
		expected, _ := json.Marshal(map[string]interface{}{
			"result": task,
		})
		assert.JSONEq(t, w.Body.String(), string(expected))
		mockService.AssertExpectations(t)
	})
}

func TestUpdateTask(t *testing.T) {
	mockService := new(mocks.Service)
	t.Run("InvalidParams", func(t *testing.T) {
		delivery := NewDelivery(mockService, &config.ServerConfig{})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBufferString(`{"name":"modified_task"}`))
		delivery.engine.ServeHTTP(w, req)

		assert.Equal(t, errcode.InvalidParams.StatusCode(), w.Code)
		expected, _ := json.Marshal(map[string]interface{}{
			"code":    errcode.InvalidParams.Code(),
			"message": errcode.InvalidParams.Message(),
		})
		assert.JSONEq(t, w.Body.String(), string(expected))
		mockService.AssertNotCalled(t, "UpdateTask")
	})
	t.Run("DuplicateRecords", func(t *testing.T) {
		task := &model.Task{
			ID:     1,
			Name:   "duplicated_name",
			Status: 1,
		}
		mockService.On("UpdateTask", task.ID, task.Name, task.Status).Return(nil, errcode.DuplicateRecords).Once()
		delivery := NewDelivery(mockService, &config.ServerConfig{})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/tasks/%d", task.ID), bytes.NewBufferString(fmt.Sprintf(`{"name":"%s", "status":%d}`, task.Name, task.Status)))
		delivery.engine.ServeHTTP(w, req)
		expected, _ := json.Marshal(map[string]interface{}{
			"code":    errcode.DuplicateRecords.Code(),
			"message": errcode.DuplicateRecords.Message(),
		})
		assert.Equal(t, errcode.DuplicateRecords.StatusCode(), w.Code)
		assert.JSONEq(t, w.Body.String(), string(expected))
		mockService.AssertExpectations(t)
	})
	t.Run("Success", func(t *testing.T) {
		task := &model.Task{
			ID:     1,
			Name:   "modified_task",
			Status: 1,
		}
		mockService.On("UpdateTask", task.ID, task.Name, task.Status).Return(task, nil).Once()

		delivery := NewDelivery(mockService, &config.ServerConfig{})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/tasks/%d", task.ID), bytes.NewBufferString(fmt.Sprintf(`{"name":"%s", "status":%d}`, task.Name, task.Status)))
		delivery.engine.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		expected, _ := json.Marshal(map[string]interface{}{
			"result": task,
		})
		assert.JSONEq(t, w.Body.String(), string(expected))
		mockService.AssertExpectations(t)
	})
}

func TestDeleteTask(t *testing.T) {
	mockService := new(mocks.Service)
	t.Run("RecordNotExists", func(t *testing.T) {
		id := uint32(1)
		mockService.On("DeleteTask", id).Return(errcode.RecordNotExists).Once()
		delivery := NewDelivery(mockService, &config.ServerConfig{})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/tasks/%d", id), nil)
		delivery.engine.ServeHTTP(w, req)

		assert.Equal(t, errcode.RecordNotExists.StatusCode(), w.Code)
		expected, _ := json.Marshal(map[string]interface{}{
			"code":    errcode.RecordNotExists.Code(),
			"message": errcode.RecordNotExists.Message(),
		})
		assert.JSONEq(t, w.Body.String(), string(expected))
		mockService.AssertExpectations(t)
	})
	t.Run("Success", func(t *testing.T) {
		id := uint32(1)
		mockService.On("DeleteTask", id).Return(nil).Once()
		delivery := NewDelivery(mockService, &config.ServerConfig{})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/tasks/%d", id), nil)
		delivery.engine.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		mockService.AssertExpectations(t)
	})
}
