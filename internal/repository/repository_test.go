package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wagaru/task/internal/errcode"
	"github.com/wagaru/task/internal/model"
)

func TestGetTasks(t *testing.T) {
	mockInMem := &inMem{
		data:  make(map[uint32]*model.Task),
		cache: make(map[string]uint32),
	}
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
	for _, task := range tasks {
		mockInMem.data[task.ID] = task
		mockInMem.cache[task.Name] = task.ID
	}

	result, err := mockInMem.GetTasks()
	assert.NoError(t, err)
	assert.Equal(t, result, tasks)
}

func TestGetTask(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		task := &model.Task{
			ID:     1,
			Name:   "task1",
			Status: 0,
		}
		mockInMem := &inMem{
			data:  map[uint32]*model.Task{task.ID: task},
			cache: map[string]uint32{task.Name: task.ID},
		}
		result, err := mockInMem.GetTask(task.ID)
		assert.NoError(t, err)
		assert.Equal(t, task, result)
	})
	t.Run("RecordNotExists", func(t *testing.T) {
		task := &model.Task{
			ID:     1,
			Name:   "task1",
			Status: 0,
		}
		mockInMem := &inMem{
			data:  map[uint32]*model.Task{task.ID: task},
			cache: map[string]uint32{task.Name: task.ID},
		}
		result, err := mockInMem.GetTask(uint32(99999))
		assert.Nil(t, result)
		assert.ErrorIs(t, err, errcode.RecordNotExists)
	})
}

func TestCreateTask(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockInMem := &inMem{
			data:  make(map[uint32]*model.Task),
			cache: make(map[string]uint32),
		}
		taskName := "newTask"
		result, err := mockInMem.CreateTask(taskName)
		assert.Nil(t, err)
		assert.Equal(t, result, &model.Task{
			ID:     1,
			Name:   taskName,
			Status: 0,
		})
	})
	t.Run("DuplicateRecords", func(t *testing.T) {
		task := &model.Task{
			ID:     1,
			Name:   "name_exists",
			Status: 0,
		}
		mockInMem := &inMem{
			data:  map[uint32]*model.Task{task.ID: task},
			cache: map[string]uint32{task.Name: task.ID},
		}
		result, err := mockInMem.CreateTask(task.Name)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, errcode.DuplicateRecords)
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		task := &model.Task{
			ID:     1,
			Name:   "task1",
			Status: 0,
		}
		mockInMem := &inMem{
			data:  map[uint32]*model.Task{task.ID: task},
			cache: map[string]uint32{task.Name: task.ID},
		}
		updatedName := "task1_rename"
		result, err := mockInMem.UpdateTask(task.ID, updatedName, task.Status)
		assert.NoError(t, err)
		assert.Equal(t, result, &model.Task{
			ID:     task.ID,
			Name:   updatedName,
			Status: task.Status,
		})
	})
	t.Run("RecordNotExists", func(t *testing.T) {
		task := &model.Task{
			ID:     1,
			Name:   "task1",
			Status: 0,
		}
		mockInMem := &inMem{
			data:  map[uint32]*model.Task{task.ID: task},
			cache: map[string]uint32{task.Name: task.ID},
		}
		result, err := mockInMem.UpdateTask(uint32(9999), "", 0)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, errcode.RecordNotExists)
	})
	t.Run("DuplicateRecords", func(t *testing.T) {
		mockInMem := &inMem{
			data:  make(map[uint32]*model.Task),
			cache: make(map[string]uint32),
		}
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
		for _, task := range tasks {
			mockInMem.data[task.ID] = task
			mockInMem.cache[task.Name] = task.ID
		}
		result, err := mockInMem.UpdateTask(tasks[0].ID, tasks[1].Name, tasks[0].Status)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, errcode.DuplicateRecords)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		task := &model.Task{
			ID:     1,
			Name:   "task1",
			Status: 0,
		}
		mockInMem := &inMem{
			data:  map[uint32]*model.Task{task.ID: task},
			cache: map[string]uint32{task.Name: task.ID},
		}
		err := mockInMem.DeleteTask(task.ID)
		assert.NoError(t, err)
		assert.Len(t, mockInMem.data, 0)
		assert.Len(t, mockInMem.cache, 0)
	})
	t.Run("RecordNotExists", func(t *testing.T) {
		mockInMem := &inMem{
			data:  make(map[uint32]*model.Task),
			cache: make(map[string]uint32),
		}
		err := mockInMem.DeleteTask(uint32(9999))
		assert.ErrorIs(t, err, errcode.RecordNotExists)
	})
}
