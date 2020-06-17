package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/likejehu/dearapril/models"
)

// AppController is interface for main logic
type AppController interface {
	SayHello()
	CreateProject(reqBody []byte) (p *models.Project, err error)
	UpdateProject(id string, reqBody []byte) (p *models.Project, err error)
	ReadProject(id string) (p *models.Project, err error)
	DeleteProject(id string) (err error)
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// ReadProject is for creating a new project
func (h *Handler) ReadProject(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectid") // getting id of project from the route
	p, err := h.App.ReadProject(projectID)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("this is ur project! ", p)
	w.WriteHeader(http.StatusOK)
}

// UpdateProject is for creating a new project
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectid") // getting id of project from the route
	reqBody, _ := ioutil.ReadAll(r.Body)
	h.App.UpdateProject(projectID, reqBody)
	fmt.Println("project updated!", projectID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteProject is for creating a new project
func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectid") // getting id of project from the route
	fmt.Println("project deleted!", projectID)
	w.WriteHeader(http.StatusNoContent)
}
