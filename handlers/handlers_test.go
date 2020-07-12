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
	error500        = errors.New("shit happens")
	testProjectZero = &models.Project{0, "this is the zEro", "zero!"}
	testProjectOne  = &models.Project{1, "this is the one", "hey you!"}
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
	taskJSON    = `{"id":1, "name":"taskOne","description":"desc1","position":1}
	`

	testCommentOne = &models.Comment{1, "this is text1", "Dec 1 13:05:05 CET 2020"}
	testCommentTwo = &models.Comment{2, "this is text2", "Dec 2 15:07:08 CET 2020"}
	testComments   = []*models.Comment{testCommentOne, testCommentTwo}
	commentJSON    = `{"id":1, "text":"this is text1","date":"Dec 1 13:05:05 CET 2020"}
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
		mockApp.On("ReadColumns", mock.Anything).Return(testColumns, nil)
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

func TestMoveColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPatch, "/", nil)
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("MoveColumn", mock.Anything, mock.Anything).Return(nil)
		handler := &Handler{&mockApp}
		handler.MoveColumn(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

//Tasks

func TestGetAllTasks(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("ReadTasks", mock.Anything).Return(testTasks, nil)
		handler := &Handler{&mockApp}
		handler.GetAllTasks(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestCreateTask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(taskJSON))
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("CreateTask", testTaskOne, mock.Anything).Return(1, nil)
		handler := &Handler{&mockApp}
		handler.CreateTask(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestReadTask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("ReadTask", mock.Anything).Return(testTaskOne, nil)
		handler := &Handler{&mockApp}
		handler.ReadTask(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(taskJSON))
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("UpdateTask", mock.Anything, testTaskOne).Return(nil)
		handler := &Handler{&mockApp}
		handler.UpdateTask(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("DeleteTask", mock.Anything).Return(nil)
		handler := &Handler{&mockApp}
		handler.DeleteTask(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}

func TestMoveTaskToColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPatch, "/", nil)
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("MoveTaskToColumn", mock.Anything, mock.Anything).Return(nil)
		handler := &Handler{&mockApp}
		handler.MoveTaskToColumn(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
func TestMoveTask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPatch, "/", nil)
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("MoveTaskUpDown", mock.Anything, mock.Anything).Return(nil)
		handler := &Handler{&mockApp}
		handler.MoveTask(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

//Comments

func TestGetAllComments(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("ReadComments", mock.Anything).Return(testComments, nil)
		handler := &Handler{&mockApp}
		handler.GetAllComments(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestCreateComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(commentJSON))
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("CreateComment", testCommentOne, mock.Anything).Return(1, nil)
		handler := &Handler{&mockApp}
		handler.CreateComment(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestReadComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("ReadComment", mock.Anything).Return(testCommentOne, nil)
		handler := &Handler{&mockApp}
		handler.ReadComment(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestUpdateComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(commentJSON))
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("UpdateComment", mock.Anything, testCommentOne).Return(nil)
		handler := &Handler{&mockApp}
		handler.UpdateComment(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestDeleteComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		mockApp := mocks.AppController{}
		mockApp.On("DeleteComment", mock.Anything).Return(nil)
		handler := &Handler{&mockApp}
		handler.DeleteComment(rec, req)
		mockApp.AssertExpectations(t)
		//assertions
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}
