package service

import (
	"github.com/wagaru/task/internal/model"
	"github.com/wagaru/task/internal/repository"
)

type service struct {
	repo repository.Repository
}

type Service interface {
	GetTasks() ([]*model.Task, error)
	GetTask(id uint32) (*model.Task, error)
	CreateTask(name string) (*model.Task, error)
	UpdateTask(id uint32, name string, status uint8) (*model.Task, error)
	DeleteTask(uint32) error
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetTasks() ([]*model.Task, error) {
	return s.repo.GetTasks()
}

func (s *service) GetTask(id uint32) (*model.Task, error) {
	return s.repo.GetTask(id)
}

func (s *service) CreateTask(name string) (*model.Task, error) {
	return s.repo.CreateTask(name)
}

func (s *service) UpdateTask(id uint32, name string, status uint8) (*model.Task, error) {
	return s.repo.UpdateTask(id, name, status)
}

func (s *service) DeleteTask(id uint32) error {
	return s.repo.DeleteTask(id)
}
