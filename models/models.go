package models

//Project must contain at least one column
type Project struct {
	Name        string
	Description string
}

// Column is manageble by user
type Column struct {
	Status string
}

// Task can be created only within a Column
type Task struct {
	Name        string
	Description string
}

// Comment can be created only within a Task
type Comment struct {
	Text string
}
