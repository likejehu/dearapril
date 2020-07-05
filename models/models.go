package models

import "time"

//Project must contain at least one column
type Project struct {
	ID          int
	Name        string
	Description string
}

// Column is manageble by user
type Column struct {
	ID       int
	Name     string
	Position string
}

// Task can be created only within a Column
type Task struct {
	ID          int
	Name        string
	Description string
	Position    string
}

// Comment can be created only within a Task
type Comment struct {
	ID   int
	Name string
	Text string
	Date time.Time
}
