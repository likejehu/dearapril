package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/likejehu/dearapril/models"
)

// AppController is interface for main logic
type AppController interface {
	SayHello()
	CreateProject(reqBody []byte) (p *models.Project, err error)
}

// Handler is struct for handlers
type Handler struct {
	App AppController
}

// Hello just says hello
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	h.App.SayHello()
	w.WriteHeader(http.StatusOK)
}

// CreateProject is for creating a new project
func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("project created!")
	fmt.Println(reqBody)
	h.App.CreateProject(reqBody)
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
