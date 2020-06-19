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
