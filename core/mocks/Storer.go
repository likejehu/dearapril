// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/likejehu/dearapril/models"
	mock "github.com/stretchr/testify/mock"
)

// Storer is an autogenerated mock type for the Storer type
type Storer struct {
	mock.Mock
}

// CreateColumn provides a mock function with given fields: column
func (_m *Storer) CreateColumn(column *models.Column) (int, error) {
	ret := _m.Called(column)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.Column) int); ok {
		r0 = rf(column)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Column) error); ok {
		r1 = rf(column)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateColumnTasks provides a mock function with given fields: taskid, colid
func (_m *Storer) CreateColumnTasks(taskid int, colid int) error {
	ret := _m.Called(taskid, colid)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(taskid, colid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateComment provides a mock function with given fields: comment
func (_m *Storer) CreateComment(comment *models.Comment) (int, error) {
	ret := _m.Called(comment)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.Comment) int); ok {
		r0 = rf(comment)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Comment) error); ok {
		r1 = rf(comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProject provides a mock function with given fields: project
func (_m *Storer) CreateProject(project *models.Project) (int, error) {
	ret := _m.Called(project)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.Project) int); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Project) error); ok {
		r1 = rf(project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProjectColumns provides a mock function with given fields: projid, colid
func (_m *Storer) CreateProjectColumns(projid int, colid int) error {
	ret := _m.Called(projid, colid)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(projid, colid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTask provides a mock function with given fields: task
func (_m *Storer) CreateTask(task *models.Task) (int, error) {
	ret := _m.Called(task)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.Task) int); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Task) error); ok {
		r1 = rf(task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTaskComments provides a mock function with given fields: taskid, comid
func (_m *Storer) CreateTaskComments(taskid int, comid int) error {
	ret := _m.Called(taskid, comid)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(taskid, comid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteColumn provides a mock function with given fields: id
func (_m *Storer) DeleteColumn(id int) error {
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
func (_m *Storer) DeleteComment(id int) error {
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
func (_m *Storer) DeleteProject(id int) error {
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
func (_m *Storer) DeleteTask(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllColumns provides a mock function with given fields: projectID
func (_m *Storer) GetAllColumns(projectID int) ([]*models.Column, error) {
	ret := _m.Called(projectID)

	var r0 []*models.Column
	if rf, ok := ret.Get(0).(func(int) []*models.Column); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Column)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllComments provides a mock function with given fields: taskID
func (_m *Storer) GetAllComments(taskID int) ([]*models.Comment, error) {
	ret := _m.Called(taskID)

	var r0 []*models.Comment
	if rf, ok := ret.Get(0).(func(int) []*models.Comment); ok {
		r0 = rf(taskID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllProjects provides a mock function with given fields:
func (_m *Storer) GetAllProjects() ([]*models.Project, error) {
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

// GetAllTasks provides a mock function with given fields: columnID
func (_m *Storer) GetAllTasks(columnID int) ([]*models.Task, error) {
	ret := _m.Called(columnID)

	var r0 []*models.Task
	if rf, ok := ret.Get(0).(func(int) []*models.Task); ok {
		r0 = rf(columnID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(columnID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetColumn provides a mock function with given fields: id
func (_m *Storer) GetColumn(id int) (*models.Column, error) {
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

// GetComment provides a mock function with given fields: id
func (_m *Storer) GetComment(id int) (*models.Comment, error) {
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

// GetProject provides a mock function with given fields: id
func (_m *Storer) GetProject(id int) (*models.Project, error) {
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

// GetTask provides a mock function with given fields: id
func (_m *Storer) GetTask(id int) (*models.Task, error) {
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

// MoveTasks provides a mock function with given fields: colid, nextid
func (_m *Storer) MoveTasks(colid int, nextid int) error {
	ret := _m.Called(colid, nextid)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(colid, nextid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateColumn provides a mock function with given fields: id, column
func (_m *Storer) UpdateColumn(id int, column *models.Column) error {
	ret := _m.Called(id, column)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.Column) error); ok {
		r0 = rf(id, column)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateColumnPosition provides a mock function with given fields: id, position
func (_m *Storer) UpdateColumnPosition(id int, position int) error {
	ret := _m.Called(id, position)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(id, position)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateColumnTasks provides a mock function with given fields: colid, taskid
func (_m *Storer) UpdateColumnTasks(colid int, taskid int) error {
	ret := _m.Called(colid, taskid)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(colid, taskid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateComment provides a mock function with given fields: id, comment
func (_m *Storer) UpdateComment(id int, comment *models.Comment) error {
	ret := _m.Called(id, comment)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.Comment) error); ok {
		r0 = rf(id, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProject provides a mock function with given fields: id, project
func (_m *Storer) UpdateProject(id int, project *models.Project) error {
	ret := _m.Called(id, project)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.Project) error); ok {
		r0 = rf(id, project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTask provides a mock function with given fields: id, task
func (_m *Storer) UpdateTask(id int, task *models.Task) error {
	ret := _m.Called(id, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *models.Task) error); ok {
		r0 = rf(id, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTaskPosition provides a mock function with given fields: id, position
func (_m *Storer) UpdateTaskPosition(id int, position int) error {
	ret := _m.Called(id, position)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(id, position)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
