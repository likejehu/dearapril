package core

import (
	"github.com/likejehu/dearapril/models"
)

//Storer is interface for database operations
type Storer interface {
	CreateProject(project *models.Project) (id int, err error)
	GetProject(id int) (project *models.Project, err error)
	GetAllProjects() (projects []*models.Project, err error)
	UpdateProject(id int, project *models.Project) (err error)
	DeleteProject(id int) (err error)

	CreateColumn(column *models.Column) (id int, err error)
	GetColumn(id int) (column *models.Column, err error)
	GetAllColumns() (columns []*models.Column, err error)
	UpdateColumn(id int, column *models.Column) (err error)
	DeleteColumn(id int) (err error)
	UpdateColumnPosition(id int, position int) (err error)

	CreateTask(task *models.Task) (id int, err error)
	GetTask(id int) (task *models.Task, err error)
	GetAllTasks() (tasks []*models.Task, err error)
	UpdateTask(id int, task *models.Task) (err error)
	DeleteTask(id int) (err error)
	UpdateTaskPosition(id int, position int) (err error)

	CreateComment(comment *models.Comment) (id int, err error)
	GetComment(id int) (comment *models.Comment, err error)
	GetAllComments() (comments []*models.Comment, err error)
	UpdateComment(id int, comment *models.Comment) (err error)
	DeleteComment(id int) (err error)

	CreateProjectColumns(projid int, colid int) (err error)
	CreateColumnTasks(taskid int, colid int) (err error)
	CreateTaskComments(taskid int, comid int) (err error)

	UpdateColumnTasks(colid int, taskid int) (err error)
	MoveTasks(colid int, nextid int) (err error)
}

//Controller is struct that has core functionality of the app
type Controller struct {
	Store Storer
}

//PROJECTS

// ReadProjects  gets all the projects
func (c *Controller) ReadProjects() (p []*models.Project, err error) {

	p, err = c.Store.GetAllProjects()
	return p, err
}

// CreateProject  creates new project
func (c *Controller) CreateProject(p *models.Project) (projid int, err error) {

	projid, _ = c.Store.CreateProject(p)
	col := new(models.Column)
	colid, _ := c.Store.CreateColumn(col)

	c.Store.CreateProjectColumns(projid, colid)
	return projid, nil
}

// UpdateProject  updates properties of  given project
func (c *Controller) UpdateProject(id int, p *models.Project) (err error) {

	c.Store.UpdateProject(id, p)
	return nil
}

// ReadProject  gets given project
func (c *Controller) ReadProject(id int) (p *models.Project, err error) {

	p, err = c.Store.GetProject(id)
	return p, err
}

// DeleteProject  deletes given project
func (c *Controller) DeleteProject(id int) (err error) {
	c.Store.DeleteProject(id)
	return nil
}

//COLUMNS

// ReadColumns  gets all the Columns
func (c *Controller) ReadColumns() (col []*models.Column, err error) {

	col, err = c.Store.GetAllColumns()
	return col, err
}

// CreateColumn  creates new Column
func (c *Controller) CreateColumn(col *models.Column, projid int) (colid int, err error) {
	colid, _ = c.Store.CreateColumn(col)
	c.Store.CreateProjectColumns(projid, colid)
	return colid, nil
}

// UpdateColumn  updates properties of  given Column
func (c *Controller) UpdateColumn(id int, col *models.Column) (err error) {

	c.Store.UpdateColumn(id, col)
	return nil
}

// ReadColumn  gets given Column
func (c *Controller) ReadColumn(id int) (col *models.Column, err error) {

	col, err = c.Store.GetColumn(id)

	return col, err
}

// DeleteColumn  deletes given Column
func (c *Controller) DeleteColumn(id int) (err error) {
	c.Store.DeleteColumn(id)
	return nil
}

// MoveColumn  moves given Column to left or right
func (c *Controller) MoveColumn(id int, next int) (err error) {
	col, err := c.Store.GetColumn(id)
	oldPosition := col.Position
	nextCol, err := c.Store.GetColumn(next)
	nextPosition := nextCol.Position
	c.Store.UpdateColumnPosition(next, oldPosition)
	if nextPosition > oldPosition {
		c.Store.UpdateColumnPosition(id, oldPosition+1)
	}
	if nextPosition > oldPosition {
		c.Store.UpdateColumnPosition(id, oldPosition-1)
	}
	return nil
}

//TASKS

// ReadTasks  gets all the Tasks
func (c *Controller) ReadTasks() (tasks []*models.Task, err error) {

	tasks, err = c.Store.GetAllTasks()
	return tasks, err
}

// CreateTask  creates new Task
func (c *Controller) CreateTask(t *models.Task, colid int) (taskid int, err error) {

	taskid, _ = c.Store.CreateTask(t)
	c.Store.CreateColumnTasks(colid, taskid)
	return taskid, nil
}

// UpdateTask  updates properties of  given Task
func (c *Controller) UpdateTask(id int, t *models.Task) (err error) {

	c.Store.UpdateTask(id, t)

	return nil
}

// ReadTask  gets given Task
func (c *Controller) ReadTask(id int) (t *models.Task, err error) {

	t, err = c.Store.GetTask(id)
	return t, err
}

// DeleteTask  deletes given Task
func (c *Controller) DeleteTask(id int) (err error) {
	c.Store.DeleteTask(id)

	return nil
}

// MoveTaskToColumn  moves given Task to left/right (across the Columns)
func (c *Controller) MoveTaskToColumn(colid, taskid int) (err error) {

	c.Store.UpdateColumnTasks(colid, taskid)
	return nil
}

// MoveTasksToColumn  moves all the task of given column to another column
func (c *Controller) MoveTasksToColumn(colid, nextid int) (err error) {

	c.Store.MoveTasks(colid, nextid)
	return nil
}

// MoveTaskUpDown  moves given Task  up/down (to prioritize it)
func (c *Controller) MoveTaskUpDown(taskid int, next int) (err error) {

	task, err := c.Store.GetTask(taskid)
	oldPosition := task.Position
	nextTask, err := c.Store.GetTask(next)
	nextPosition := nextTask.Position
	c.Store.UpdateTaskPosition(next, oldPosition)
	if nextPosition > oldPosition {
		c.Store.UpdateTaskPosition(taskid, oldPosition+1)
	}
	if nextPosition > oldPosition {
		c.Store.UpdateTaskPosition(taskid, oldPosition-1)
	}
	return nil
}

//COMMENTS

// ReadComments  gets all the Comments
func (c *Controller) ReadComments() (com []*models.Comment, err error) {

	com, err = c.Store.GetAllComments()
	return com, err
}

// CreateComment  creates new Comment
func (c *Controller) CreateComment(com *models.Comment, taskid int) (comid int, err error) {
	comid, _ = c.Store.CreateComment(com)
	c.Store.CreateTaskComments(taskid, comid)
	return comid, nil
}

// UpdateComment   updates properties of  given Comment
func (c *Controller) UpdateComment(id int, com *models.Comment) (err error) {

	c.Store.UpdateComment(id, com)
	return nil
}

// ReadComment   gets given Comment
func (c *Controller) ReadComment(id int) (com *models.Comment, err error) {
	com, err = c.Store.GetComment(id)
	return com, err
}

// DeleteComment   deletes given Comment
func (c *Controller) DeleteComment(id int) (err error) {
	err = c.Store.DeleteColumn(id)

	return err
}
