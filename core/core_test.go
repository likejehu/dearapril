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
