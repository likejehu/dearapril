package core

import (
	"fmt"

	"github.com/likejehu/dearapril/models"
)

//Controller is struct that has core functionality of the app
type Controller struct {
}

// SayHello just says hello
func (c *Controller) SayHello() {
	fmt.Println("hello!")
}

// CreateProject  creates new project
func (c *Controller) CreateProject() (p *models.Project, err error) {
	p = new(models.Project)
	return p, nil
}

// UpdateProject  updates properties of  given project
func (c *Controller) UpdateProject(id string) (p *models.Project, err error) {
	p = new(models.Project)
	return p, nil
}

// ReadProject  gets given project
func (c *Controller) ReadProject(id string) (p *models.Project, err error) {
	p = new(models.Project)
	return p, nil
}

// DeleteProject  deletes given project
func (c *Controller) DeleteProject(id string) (err error) {

	return nil
}

// App is instance of Controller
var App = new(Controller)
