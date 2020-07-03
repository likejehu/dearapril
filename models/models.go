package models

import "time"

//Project must contain at least one column
type Project struct {
	ID          int
	Name        string
	Description string
	Columns
}

// Column is manageble by user
type Column struct {
	ID       int
	Name     string
	Position int
	Tasks
}

// Columns is struct to store columns within map
type Columns struct {
	Cls map[int]*Column
}

// Task can be created only within a Column
type Task struct {
	ID          int
	Name        string
	Description string
	Position    int
	Comments
}

// Tasks is struct to store tasks within map
type Tasks struct {
	Tks map[int]*Task
}

// Comment can be created only within a Task
type Comment struct {
	ID   int
	Name string
	Text string
	Date time.Time
}

// Comments is struct to store comments within map
type Comments struct {
	Cms map[int]*Comment
}
