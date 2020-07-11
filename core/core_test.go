package core

import (
	"errors"
	"testing"

	"github.com/likejehu/dearapril/core/mocks"
	"github.com/likejehu/dearapril/models"
	"github.com/stretchr/testify/assert"
)

var (
	error500        = errors.New("shit happens")
	testProjectZero = &models.Project{0, "this is the zEro", "zero!"}
	testProjectOne  = &models.Project{1, "this is the one", "hey you!"}
	testProjectTwo  = &models.Project{2, "this is the two", "how are you doin?"}
	testProjects    = []*models.Project{testProjectOne, testProjectTwo}

	testColumnOne  = &models.Column{1, "columnOne", 1}
	testColumnTwo  = &models.Column{2, "columnTwo", 2}
	testColumnZero = &models.Column{ID: 0, Name: "", Position: 0}
	testColumns    = []*models.Column{testColumnOne, testColumnTwo}

	testTaskOne = &models.Task{1, "taskOne", "desc1", 1}
	testTaskTwo = &models.Task{1, "taskTwo", "desc2", 2}
	testTasks   = []*models.Task{testTaskOne, testTaskTwo}

	testCommentOne = &models.Comment{1, "this is text1", "Dec 1 13:05:05 CET 2020"}
	testCommentTwo = &models.Comment{2, "this is text2", "Dec 2 15:07:08 CET 2020"}
	testComments   = []*models.Comment{testCommentOne, testCommentTwo}
)

//PROJECTS

func TestReadProjects(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("GetAllProjects").Return(testProjects, nil)
		projects, _ := testController.ReadProjects()
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, testProjects, projects)
	})
}

func TestCreateProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("CreateProject", testProjectOne).Return(1, nil)
		mockStore.On("CreateColumn", testColumnZero).Return(1, nil)
		mockStore.On("CreateProjectColumns", 1, 1).Return(nil)
		projectid, _ := testController.CreateProject(testProjectOne)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, projectid, 1)
	})
}
