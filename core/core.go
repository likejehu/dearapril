package core

import (
	"encoding/json"
	"fmt"

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

	CreateTask(task *models.Task) (id int, err error)
	GetTask(id int) (task *models.Task, err error)
	GetAllTasks() (tasks []*models.Task, err error)
	UpdateTask(id int, task *models.Task) (err error)
	DeleteTask(id int) (err error)

	CreateComment(comment *models.Comment) (id int, err error)
	GetComment(id int) (comment *models.Comment, err error)
	GetAllComments() (comments []*models.Comment, err error)
	UpdateComment(id int, comment *models.Comment) (err error)
	DeleteComment(id int) (err error)

	CreateProjectColumns(projid int, colid int) (err error)
	CreateColumnTasks(taskid int, colid int) (err error)
	CreateTaskComments(taskid int, comid int) (err error)

	UpdateColumnTasks(colid int, taskid int) (err error)
}

//Controller is struct that has core functionality of the app
type Controller struct {
	Store Storer
}

// SayHello just says hello
func (c *Controller) SayHello() {
	fmt.Println("hello!")
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

// CreateColumn  creates new Column
func (c *Controller) CreateColumn(col *models.Column, projid int) (colid int, err error) {
	colid, _ = c.Store.CreateColumn(col)
	c.Store.CreateProjectColumns(projid, colid)
	return colid, nil
}

// UpdateColumn  updates properties of  given Column
func (c *Controller) UpdateColumn(id int, col *models.Column) (colid int, err error) {

	c.Store.UpdateColumn(id, col)
	return colid, nil
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
func (c *Controller) MoveColumn(id int, reqBody []byte) (err error) {
	col := new(models.Column)
	json.Unmarshal(reqBody, &col)
	c.Store.UpdateColumn(id, col)

	return nil
}

//TASKS

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

// MoveTaskToColumn  moves given Task to left/right (across the Columns) or up/down (to prioritize it)
func (c *Controller) MoveTaskToColumn(colid, taskid int, reqBody []byte) (err error) {

	c.Store.UpdateColumnTasks(colid, taskid)
	return nil
}

// MoveTaskUpDown  moves given Task to left/right (across the Columns) or up/down (to prioritize it)
func (c *Controller) MoveTaskUpDown(taskid int, t *models.Task) (err error) {

	c.Store.UpdateTask(taskid, t)
	return nil
}

//COMMENTS

// CreateComment  creates new Task
func (c *Controller) CreateComment(reqBody []byte) (com *models.Comment, err error) {
	com = new(models.Comment)
	json.Unmarshal(reqBody, &c)
	c.Store.CreateComment(com)
	return com, nil
}

// UpdateComment   updates properties of  given Comment
func (c *Controller) UpdateComment(id int, reqBody []byte) (err error) {
	com := new(models.Comment)

	json.Unmarshal(reqBody, &com)
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
