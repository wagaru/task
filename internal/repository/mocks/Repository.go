// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/wagaru/task/internal/model"

	testing "testing"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: name
func (_m *Repository) CreateTask(name string) (*model.Task, error) {
	ret := _m.Called(name)

	var r0 *model.Task
	if rf, ok := ret.Get(0).(func(string) *model.Task); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTask provides a mock function with given fields: id
func (_m *Repository) DeleteTask(id uint32) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint32) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTask provides a mock function with given fields: id
func (_m *Repository) GetTask(id uint32) (*model.Task, error) {
	ret := _m.Called(id)

	var r0 *model.Task
	if rf, ok := ret.Get(0).(func(uint32) *model.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasks provides a mock function with given fields:
func (_m *Repository) GetTasks() ([]*model.Task, error) {
	ret := _m.Called()

	var r0 []*model.Task
	if rf, ok := ret.Get(0).(func() []*model.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: id, name, status
func (_m *Repository) UpdateTask(id uint32, name string, status uint8) (*model.Task, error) {
	ret := _m.Called(id, name, status)

	var r0 *model.Task
	if rf, ok := ret.Get(0).(func(uint32, string, uint8) *model.Task); ok {
		r0 = rf(id, name, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32, string, uint8) error); ok {
		r1 = rf(id, name, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a cleanup function to assert the mocks expectations.
func NewRepository(t testing.TB) *Repository {
	mock := &Repository{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
