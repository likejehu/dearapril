// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/likejehu/dearapril/models"
	mock "github.com/stretchr/testify/mock"
)

// AppController is an autogenerated mock type for the AppController type
type AppController struct {
	mock.Mock
}

// CreateColumn provides a mock function with given fields: col, projid
func (_m *AppController) CreateColumn(col *models.Column, projid int) (int, error) {
	ret := _m.Called(col, projid)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.Column, int) int); ok {
		r0 = rf(col, projid)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Column, int) error); ok {
		r1 = rf(col, projid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateComment provides a mock function with given fields: com, taskid
func (_m *AppController) CreateComment(com *models.Comment, taskid int) (int, error) {
	ret := _m.Called(com, taskid)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.Comment, int) int); ok {
		r0 = rf(com, taskid)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Comment, int) error); ok {
		r1 = rf(com, taskid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProject provides a mock function with given fields: p
func (_m *AppController) CreateProject(p *models.Project) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.Project) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Project) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTask provides a mock function with given fields: t, colid
func (_m *AppController) CreateTask(t *models.Task, colid int) (int, error) {
	ret := _m.Called(t, colid)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.Task, int) int); ok {
		r0 = rf(t, colid)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Task, int) error); ok {
		r1 = rf(t, colid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteColumn provides a mock function with given fields: id
func (_m *AppController) DeleteColumn(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComment provides a mock function with given fields: id
func (_m *AppController) DeleteComment(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProject provides a mock function with given fields: id
func (_m *AppController) DeleteProject(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTask provides a mock function with given fields: id
func (_m *AppController) DeleteTask(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MoveColumn provides a mock function with given fields: id, next
func (_m *AppController) MoveColumn(id int, next int) error {
	ret := _m.Called(id, next)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(id, next)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MoveTaskToColumn provides a mock function with given fields: colid, taskid
func (_m *AppController) MoveTaskToColumn(colid int, taskid int) error {
	ret := _m.Called(colid, taskid)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(colid, taskid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MoveTaskUpDown provides a mock function with given fields: taskid, next
func (_m *AppController) MoveTaskUpDown(taskid int, next int) error {
	ret := _m.Called(taskid, next)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(taskid, next)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MoveTasksToColumn provides a mock function with given fields: colid, nextid
func (_m *AppController) MoveTasksToColumn(colid int, nextid int) error {
	ret := _m.Called(colid, nextid)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(colid, nextid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadColumn provides a mock function with given fields: id
func (_m *AppController) ReadColumn(id int) (*models.Column, error) {
	ret := _m.Called(id)

	var r0 *models.Column
	if rf, ok := ret.Get(0).(func(int) *models.Column); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Column)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadColumns provides a mock function with given fields:
func (_m *AppController) ReadColumns() ([]*models.Column, error) {
	ret := _m.Called()

	var r0 []*models.Column
	if rf, ok := ret.Get(0).(func() []*models.Column); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Column)
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

// ReadComment provides a mock function with given fields: id
func (_m *AppController) ReadComment(id int) (*models.Comment, error) {
	ret := _m.Called(id)

	var r0 *models.Comment
	if rf, ok := ret.Get(0).(func(int) *models.Comment); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadComments provides a mock function with given fields:
func (_m *AppController) ReadComments() ([]*models.Comment, error) {
	ret := _m.Called()

	var r0 []*models.Comment
	if rf, ok := ret.Get(0).(func() []*models.Comment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Comment)
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

// ReadProject provides a mock function with given fields: id
func (_m *AppController) ReadProject(id int) (*models.Project, error) {
	ret := _m.Called(id)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(int) *models.Project); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadProjects provides a mock function with given fields:
func (_m *AppController) ReadProjects() ([]*models.Project, error) {
	ret := _m.Called()

	var r0 []*models.Project
	if rf, ok := ret.Get(0).(func() []*models.Project); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Project)
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

// ReadTask provides a mock function with given fields: id
func (_m *AppController) ReadTask(id int) (*models.Task, error) {
	ret := _m.Called(id)

	var r0 *models.Task
	if rf, ok := ret.Get(0).(func(int) *models.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadTasks provides a mock function with given fields:
func (_m *AppController) ReadTasks() ([]*models.Task, error) {
	ret := _m.Called()

	var r0 []*models.Task
	if rf, ok := ret.Get(0).(func() []*models.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Task)
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

// UpdateColumn provides a mock function with given fields: id, col
func (_m *AppController) UpdateColumn(id int, col *models.Column) error {
	ret := _m.Called(id, col)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.Column) error); ok {
		r0 = rf(id, col)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateComment provides a mock function with given fields: id, com
func (_m *AppController) UpdateComment(id int, com *models.Comment) error {
	ret := _m.Called(id, com)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.Comment) error); ok {
		r0 = rf(id, com)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProject provides a mock function with given fields: id, p
func (_m *AppController) UpdateProject(id int, p *models.Project) error {
	ret := _m.Called(id, p)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.Project) error); ok {
		r0 = rf(id, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTask provides a mock function with given fields: id, t
func (_m *AppController) UpdateTask(id int, t *models.Task) error {
	ret := _m.Called(id, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.Task) error); ok {
		r0 = rf(id, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
