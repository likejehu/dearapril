package models

//Project must contain at least one column
type Project struct {
	ID          int
	Name        string
	Description string
	Columns
}

// Column is manageble by user
type Column struct {
	Status string
	Tasks
}

// Columns is struct to store columns within map
type Columns struct {
	Cls map[string]*Column
}

// Task can be created only within a Column
type Task struct {
	Name        string
	Description string
	Comments
}

// Tasks is struct to store tasks within map
type Tasks struct {
	Tks map[string]*Task
}

// Comment can be created only within a Task
type Comment struct {
	Text string
}

// Comments is struct to store comments within map
type Comments struct {
	Cms map[string]*Comment
}
