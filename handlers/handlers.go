package handlers

import (
	"fmt"
	"net/http"

	"github.com/likejehu/dearapril/core"
)

// Handler is struct for handlers
type Handler struct {
}

// Hello just says hello
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	core.App.SayHello()
	w.WriteHeader(http.StatusOK)
}

// CreateProject is for creating a new project
func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("project created!")
	w.WriteHeader(http.StatusOK)
}

// ReadProject is for creating a new project
func (h *Handler) ReadProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is ur project!")
	w.WriteHeader(http.StatusOK)
}

// UpdateProject is for creating a new project
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("project updated!")
	w.WriteHeader(http.StatusOK)
}

// DeleteProject is for creating a new project
func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("project deleted!")
	w.WriteHeader(http.StatusNoContent)
}
