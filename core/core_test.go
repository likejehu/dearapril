package core

import (
	"errors"
	"testing"

	"github.com/likejehu/dearapril/core/mocks"
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
		mockStore.On("GetAllColumns", mock.Anything).Return(testColumns, nil)
		columns, _ := testController.ReadColumns(1)
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

//TASKS

func TestReadTasks(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("GetAllTasks", mock.Anything).Return(testTasks, nil)
		tasks, _ := testController.ReadTasks(1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, testTasks, tasks)
	})
}
func TestCreateTask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("CreateTask", testTaskOne).Return(1, nil)
		mockStore.On("CreateColumnTasks", 1, 1).Return(nil)
		taskid, _ := testController.CreateTask(testTaskOne, 1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, taskid, 1)
	})
}
func TestUpdateTask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("UpdateTask", 1, testTaskOne).Return(nil)
		testController.UpdateTask(1, testTaskOne)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}

func TestReadTask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("GetTask", 1).Return(testTaskOne, nil)
		task, _ := testController.ReadTask(1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, testTaskOne, task)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("DeleteTask", 1).Return(nil)
		testController.DeleteTask(1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}

func TestMoveTaskToColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("UpdateColumnTasks", 1, 1).Return(nil)
		testController.MoveTaskToColumn(1, 1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}

func TestMoveTasksToColumn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("MoveTasks", 1, 1).Return(nil)
		testController.MoveTasksToColumn(1, 1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}

func TestMoveTaskUpDown(t *testing.T) {
	t.Run("success with next bigger than current", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("GetTask", 1).Return(testTaskOne, nil)
		mockStore.On("GetTask", 2).Return(testTaskTwo, nil)
		mockStore.On("UpdateTaskPosition", 2, 1).Return(nil)
		mockStore.On("UpdateTaskPosition", 1, 2).Return(nil)
		testController.MoveTaskUpDown(1, 2)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}

//COMMENTS

func TestReadComments(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("GetAllComments", mock.Anything).Return(testComments, nil)
		comments, _ := testController.ReadComments(1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, testComments, comments)
	})
}

func TestCreateComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("CreateComment", testCommentOne).Return(1, nil)
		mockStore.On("CreateTaskComments", 1, 1).Return(nil)
		commentid, _ := testController.CreateComment(testCommentOne, 1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, commentid, 1)
	})
}

func TestUpdateComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("UpdateComment", 1, testCommentOne).Return(nil)
		testController.UpdateComment(1, testCommentOne)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}

func TestReadComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("GetComment", 1).Return(testCommentOne, nil)
		comment, _ := testController.ReadComment(1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, testCommentOne, comment)
	})
}

func TestDeleteComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		mockStore := mocks.Storer{}
		testController := Controller{&mockStore}
		mockStore.On("DeleteComment", 1).Return(nil)
		testController.DeleteComment(1)
		mockStore.AssertExpectations(t)
		//assertions
		assert.Equal(t, nil, nil)
	})
}
