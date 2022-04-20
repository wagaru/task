package repository

import (
	"sync"
	"sync/atomic"

	"github.com/wagaru/task/internal/errcode"
	"github.com/wagaru/task/internal/model"
)

var UID uint32 = 0

type inMem struct {
	mux   sync.RWMutex
	data  map[uint32]*model.Task
	cache map[string]uint32
}

type Repository interface {
	GetTasks() ([]*model.Task, error)
	GetTask(id uint32) (*model.Task, error)
	CreateTask(name string) (*model.Task, error)
	UpdateTask(id uint32, name string, status uint8) (*model.Task, error)
	DeleteTask(id uint32) error
}

func NewRepository() Repository {
	return &inMem{
		data:  make(map[uint32]*model.Task),
		cache: make(map[string]uint32),
	}
}

func (in *inMem) GetTasks() ([]*model.Task, error) {
	in.mux.RLock()
	defer in.mux.RUnlock()
	tasks := make([]*model.Task, 0, len(in.data))
	for _, task := range in.data {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (in *inMem) GetTask(id uint32) (*model.Task, error) {
	in.mux.RLock()
	defer in.mux.RUnlock()
	task, ok := in.data[id]
	if !ok {
		return nil, errcode.RecordNotExists
	}
	return task, nil
}

func (in *inMem) CreateTask(name string) (*model.Task, error) {
	in.mux.Lock()
	defer in.mux.Unlock()

	if _, ok := in.cache[name]; ok {
		return nil, errcode.DuplicateRecords
	}
	task := &model.Task{
		Name:   name,
		ID:     atomic.AddUint32(&UID, 1),
		Status: 0,
	}

	in.data[task.ID] = task
	in.cache[task.Name] = task.ID
	return task, nil
}

func (in *inMem) UpdateTask(id uint32, name string, status uint8) (*model.Task, error) {
	in.mux.Lock()
	defer in.mux.Unlock()

	task, ok := in.data[id]
	if !ok {
		return nil, errcode.RecordNotExists
	}

	if cid, ok := in.cache[name]; ok && cid != id {
		return nil, errcode.DuplicateRecords
	}
	task.Name = name
	task.Status = status

	in.data[task.ID] = task
	in.cache[task.Name] = task.ID
	return task, nil
}

func (in *inMem) DeleteTask(id uint32) error {
	in.mux.Lock()
	defer in.mux.Unlock()

	task, ok := in.data[id]
	if !ok {
		return errcode.RecordNotExists
	}

	delete(in.data, id)
	delete(in.cache, task.Name)
	return nil
}
