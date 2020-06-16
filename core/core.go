package core

import (
	"encoding/json"
	"fmt"

	"github.com/likejehu/dearapril/models"
)

//Storer is interface for database operations
type Storer interface {
	Post(id string)
	Get(id string)
	Update(id string)
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
	return p, nil
}

// UpdateProject  updates properties of  given project
func (c *Controller) UpdateProject(id string) (p *models.Project, err error) {
	p = new(models.Project)

	c.Store.Update(id)
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
