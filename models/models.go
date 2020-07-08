package models

import "time"

//Project must contain at least one column
type Project struct {
	ID          int
	Name        string `json:"name"  validate:"required"`
	Description string `json:"description"  validate:"required"`
}

// Column is manageble by user
type Column struct {
	ID       int
	Name     string `json:"name"  validate:"required"`
	Position string
}

// Task can be created only within a Column
type Task struct {
	ID          int
	Name        string `json:"name"  validate:"required"`
	Description string `json:"description"  validate:"required"`
	Position    string
}

// Comment can be created only within a Task
type Comment struct {
	ID   int
	Name string `json:"name"  validate:"required"`
	Text string `json:"text"  validate:"required"`
	Date time.Time
}
