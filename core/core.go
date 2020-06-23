package core

import (
	"encoding/json"
	"fmt"

	"github.com/likejehu/dearapril/models"
)

//Storer is interface for database operations
type Storer interface {
	Post(p *models.Project)
	Get(id string)
	GetAll()
	Update(id string, p *models.Project)
	Delete(id string)
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
func (c *Controller) CreateColumn(reqBody []byte) (p *models.Project, err error) {
	p = new(models.Project)
	json.Unmarshal(reqBody, &p)
	c.Store.Post(p)
	return p, nil
}

// UpdateColumn  updates properties of  given Column
func (c *Controller) UpdateColumn(id string, reqBody []byte) (p *models.Project, err error) {
	p = new(models.Project)

	json.Unmarshal(reqBody, &p)
	c.Store.Update(id, p)
	return p, nil
}

// ReadColumn  gets given pColumn
func (c *Controller) ReadColumn(id string) (p *models.Project, err error) {
	p = new(models.Project)
	c.Store.Get(id)
	return p, nil
}

// DeleteColumn  deletes given Column
func (c *Controller) DeleteColumn(id string) (err error) {
	c.Store.Delete(id)

	return nil
}

// RenameColumn  renames given Column
func (c *Controller) RenameColumn(id string) (err error) {

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
