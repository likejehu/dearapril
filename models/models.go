package models

import "time"

//Project must contain at least one column
type Project struct {
	ID          int
	Name        string `json:"name"  validate:"required,min=1,max=500"`
	Description string `json:"description"  validate:"required,min=0,max=1000"`
}

// Column is manageble by user
type Column struct {
	ID       int
	Name     string `json:"name"  validate:"required,min=1,max=255"`
	Position int
}

// Task can be created only within a Column
type Task struct {
	ID          int
	Name        string `json:"name"  validate:"required,min=1,max=500"`
	Description string `json:"description"  validate:"required,min=0,max=5000"`
	Position    int
}

// Comment can be created only within a Task
type Comment struct {
	ID   int
	Text string `json:"text"  validate:"required,min=1,max=5000"`
	Date time.Time
}
