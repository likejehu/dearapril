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
)

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
