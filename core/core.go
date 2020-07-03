package core

import (
	"encoding/json"
	"fmt"

	"github.com/likejehu/dearapril/models"
)

//Storer is interface for database operations
type Storer interface {
	CreateProject(project *models.Project) (err error)
	GetProject(id int) (project *models.Project, err error)
	GetAllProjects() (projects []*models.Project, err error)
	UpdateProject(id int, project *models.Project) (err error)
	DeleteProject(id int) (err error)

	CreateColumn(column *models.Column) (err error)
	GetColumn(id int) (column *models.Column, err error)
	GetAllColumns() (columns []*models.Column, err error)
	UpdateColumn(id int, column *models.Column) (err error)
	DeleteColumn(id int) (err error)

	CreateTask(task *models.Task) (err error)
	GetTask(id int) (task *models.Task, err error)
	GetAllTasks() (tasks []*models.Task, err error)
	UpdateTask(id int, task *models.Task) (err error)
	DeleteTask(id int) (err error)

	CreateComment(comment *models.Comment) (err error)
	GetComment(id int) (comment *models.Comment, err error)
	GetAllComments() (comments []*models.Comment, err error)
	UpdateComment(id int, comment *models.Comment) (err error)
	DeleteComment(id int) (err error)
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
func (c *Controller) ReadProjects() (p *models.Project, err error) {
	p = new(models.Project)
	c.Store.GetAll()
	return p, nil
}

// CreateProject  creates new project
func (c *Controller) CreateProject(reqBody []byte) (p *models.Project, err error) {
	p = new(models.Project)
	json.Unmarshal(reqBody, &p)
	c.Store.Post(p)
	return p, nil
}

// UpdateProject  updates properties of  given project
func (c *Controller) UpdateProject(id string, reqBody []byte) (p *models.Project, err error) {
	p = new(models.Project)
	json.Unmarshal(reqBody, &p)
	c.Store.Update(id, p)
	return p, nil
}

// ReadProject  gets given project
func (c *Controller) ReadProject(id string) (p *models.Project, err error) {
	p = new(models.Project)
	c.Store.Get(id)
	return p, nil
}

// DeleteProject  deletes given project
func (c *Controller) DeleteProject(id string) (err error) {
	c.Store.Delete(id)
	return nil
}

//COLUMNS

// CreateColumn  creates new Column
func (c *Controller) CreateColumn(reqBody []byte) (col *models.Column, err error) {
	proj := new(models.Project)
	col = new(models.Column)
	id := "2"
	proj.Cls[id] = col
	json.Unmarshal(reqBody, &proj)
	c.Store.Post(proj)
	return proj.Cls[id], nil
}

// UpdateColumn  updates properties of  given Column
func (c *Controller) UpdateColumn(id string, reqBody []byte) (col *models.Column, err error) {
	proj := new(models.Project)
	col = new(models.Column)
	proj.Cls[id] = col
	json.Unmarshal(reqBody, &proj)
	c.Store.Update(id, proj)
	return col, nil
}

// ReadColumn  gets given Column
func (c *Controller) ReadColumn(id string) (col *models.Column, err error) {
	proj := new(models.Project)
	c.Store.Get(id)
	col = proj.Cls[id]
	return col, nil
}

// DeleteColumn  deletes given Column
func (c *Controller) DeleteColumn(id string) (err error) {
	c.Store.Delete(id)

	return nil
}

// MoveColumn  moves given Column to left or right
func (c *Controller) MoveColumn(id string, direction string) (err error) {

	return nil
}

//TASKS

// CreateTask  creates new Task
func (c *Controller) CreateTask(reqBody []byte) (p *models.Project, err error) {
	p = new(models.Project)
	json.Unmarshal(reqBody, &p)
	c.Store.Post(p)
	return p, nil
}

// UpdateTask  updates properties of  given Task
func (c *Controller) UpdateTask(id string, reqBody []byte) (p *models.Project, err error) {
	p = new(models.Project)

	json.Unmarshal(reqBody, &p)
	c.Store.Update(id, p)
	return p, nil
}

// ReadTask  gets given Task
func (c *Controller) ReadTask(id string) (p *models.Project, err error) {
	p = new(models.Project)
	c.Store.Get(id)
	return p, nil
}

// DeleteTask  deletes given Task
func (c *Controller) DeleteTask(id string) (err error) {
	c.Store.Delete(id)

	return nil
}

// MoveTask  moves given Task to left/right (across the Columns) or up/down (to prioritize it)
func (c *Controller) MoveTask(id string, direction string) (err error) {

	return nil
}

//COMMENTS

// CreateComment  creates new Task
func (c *Controller) CreateComment(reqBody []byte) (p *models.Project, err error) {
	p = new(models.Project)
	json.Unmarshal(reqBody, &p)
	c.Store.Post(p)
	return p, nil
}

// UpdateComment   updates properties of  given Comment
func (c *Controller) UpdateComment(id string, reqBody []byte) (p *models.Project, err error) {
	p = new(models.Project)

	json.Unmarshal(reqBody, &p)
	c.Store.Update(id, p)
	return p, nil
}

// ReadComment   gets given Comment
func (c *Controller) ReadComment(id string) (p *models.Project, err error) {
	p = new(models.Project)
	c.Store.Get(id)
	return p, nil
}

// DeleteComment   deletes given Comment
func (c *Controller) DeleteComment(id string) (err error) {
	c.Store.Delete(id)

	return nil
}
