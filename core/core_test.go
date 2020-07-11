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

func TestUpdateProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("UpdateProject", 1, testProjectOne).Return(nil)
		testController.UpdateProject(1, testProjectOne)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}

func TestReadProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("GetProject", 1).Return(testProjectOne, nil)
		project, _ := testController.ReadProject(1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, testProjectOne, project)
	})
}
func TestDeleteProject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("DeleteProject", 1).Return(nil)
		testController.DeleteProject(1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}

//COLUMNS

func TestReadColumns(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("GetAllColumns").Return(testColumns, nil)
		columns, _ := testController.ReadColumns()
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, testColumns, columns)
	})
}

func TestCreateColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("CreateColumn", testColumnOne).Return(1, nil)
		mockStore.On("CreateProjectColumns", 1, 1).Return(nil)
		columnid, _ := testController.CreateColumn(testColumnOne, 1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, columnid, 1)
	})
}

func TestUpdateColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("UpdateColumn", 1, testColumnOne).Return(nil)
		testController.UpdateColumn(1, testColumnOne)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}

func TestReadColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("GetColumn", 1).Return(testColumnOne, nil)
		column, _ := testController.ReadColumn(1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, testColumnOne, column)
	})
}

func TestDeleteColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("DeleteColumn", 1).Return(nil)
		testController.DeleteColumn(1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}

func TestMoveColumn(t *testing.T) {
	t.Run("success with next bigger than current", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("GetColumn", 1).Return(testColumnOne, nil)
		mockStore.On("GetColumn", 2).Return(testColumnTwo, nil)
		mockStore.On("UpdateColumnPosition", 2, 1).Return(nil)
		mockStore.On("UpdateColumnPosition", 1, 2).Return(nil)
		testController.MoveColumn(1, 2)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}
