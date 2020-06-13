package handlers

import (
	"fmt"
	"net/http"
)

// Handler is struct for handlers
type Handler struct {
}

// Hello just says hello
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello!")
}

// CreateProject is for creating a new project
func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("project created!")
}

// ReadProject is for creating a new project
func (h *Handler) ReadProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is ur project!")
}

// UpdateProject is for creating a new project
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("project updated!")
}

// DeleteProject is for creating a new project
func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("project deleted!")
}
