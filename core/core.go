package core

import "fmt"

//Controller is struct that has core functionality of the app
type Controller struct {
}

// SayHello just says hello
func (c *Controller) SayHello() {
	fmt.Println("hello!")
}

// App is instance of Controller
var App = new(Controller)
