package models

//Project must contain at least one column
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"  validate:"required,min=1,max=500"`
	Description string `json:"description"  validate:"required,min=0,max=1000"`
}

// Column is manageble by user
type Column struct {
	ID       int    `json:"id"`
	Name     string `json:"name"  validate:"required,min=1,max=255"`
	Position int    `json:"position"`
}

// Task can be created only within a Column
type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"  validate:"required,min=1,max=500"`
	Description string `json:"description"  validate:"required,min=0,max=5000"`
	Position    int    `json:"position"`
}

// Comment can be created only within a Task
type Comment struct {
	ID   int    `json:"id"`
	Text string `json:"text"  validate:"required,min=1,max=5000"`
	Date string `json:"date"`
}
