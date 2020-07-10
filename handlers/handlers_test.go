package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/likejehu/dearapril/handlers/mocks"
	"github.com/likejehu/dearapril/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	error500 = errors.New("shit happens")

	testProjectOne  = &models.Project{1, "this is the one", "hey you!"}
	testProjectZero = &models.Project{0, "this is the zEro", "zero!"}
	testProjectTwo  = &models.Project{2, "this is the two", "how are you doin?"}
	testProjects    = []*models.Project{testProjectOne, testProjectTwo}

	projectJSON = `{"name":"this is the zEro","description":"zero!"}
	`

	testColumnOne = &models.Column{1, "columnOne", 1}
	testColumnTwo = &models.Column{2, "columnTwo", 2}
	testColumns   = []*models.Column{testColumnOne, testColumnTwo}

	columnJSON = `{"id":1, "name":"columnOne","position":1}
	`
	testTaskOne = &models.Task{1, "taskOne", "desc1", 1}
	testTaskTwo = &models.Task{1, "taskTwo", "desc2", 2}
	testTasks   = []*models.Task{testTaskOne, testTaskTwo}
	taskJSON    = `{"id":1, "name":"columnOne","description":"desc1","position":1}
	`
)

//Projects
func TestGetAllProjects(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("ReadProjects").Return(testProjects, nil)
		handler := &Handler{&mockApp}
		handler.GetAllProjects(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestCreateProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(projectJSON))
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("CreateProject", testProjectZero).Return(1, nil)
		handler := &Handler{&mockApp}
		handler.CreateProject(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestReadProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("ReadProject", mock.Anything).Return(testProjectOne, nil)
		handler := &Handler{&mockApp}
		handler.ReadProject(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestUpdateProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(projectJSON))
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("UpdateProject", mock.Anything, testProjectZero).Return(nil)
		handler := &Handler{&mockApp}
		handler.UpdateProject(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestDeleteProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("DeleteProject", mock.Anything).Return(nil)
		handler := &Handler{&mockApp}
		handler.DeleteProject(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}

//Columns

func TestGetAllColumns(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("ReadColumns").Return(testColumns, nil)
		handler := &Handler{&mockApp}
		handler.GetAllColumns(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestCreateColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(columnJSON))
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("CreateColumn", testColumnOne, mock.Anything).Return(1, nil)
		handler := &Handler{&mockApp}
		handler.CreateColumn(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestReadColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("ReadColumn", mock.Anything).Return(testColumnOne, nil)
		handler := &Handler{&mockApp}
		handler.ReadColumn(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestUpdateColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(columnJSON))
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("UpdateColumn", mock.Anything, testColumnOne).Return(nil)
		handler := &Handler{&mockApp}
		handler.UpdateColumn(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestDeleteColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("MoveTasksToColumn", mock.Anything, mock.Anything).Return(nil)
		mockApp.On("DeleteColumn", mock.Anything).Return(nil)
		handler := &Handler{&mockApp}
		handler.DeleteColumn(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}

//Tasks

func TestGetAllTasks(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("ReadTasks").Return(testTasks, nil)
		handler := &Handler{&mockApp}
		handler.GetAllTasks(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

/*
func TestCreateProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(projectJSON))
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("CreateProject", testProjectZero).Return(1, nil)
		handler := &Handler{&mockApp}
		handler.CreateProject(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestReadProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("ReadProject", mock.Anything).Return(testProjectOne, nil)
		handler := &Handler{&mockApp}
		handler.ReadProject(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestUpdateProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(projectJSON))
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("UpdateProject", mock.Anything, testProjectZero).Return(nil)
		handler := &Handler{&mockApp}
		handler.UpdateProject(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestDeleteProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		mockApp.On("DeleteProject", mock.Anything).Return(nil)
		handler := &Handler{&mockApp}
		handler.DeleteProject(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}
*/
