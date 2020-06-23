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

// GetAllProjects just says hello
func (h *Handler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
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

// ReadProject is for reading a new project
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

// UpdateProject is for updating a new project
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectid") // getting id of project from the route
	reqBody, _ := ioutil.ReadAll(r.Body)
	h.App.UpdateProject(projectID, reqBody)
	fmt.Println("project updated!", projectID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteProject is for deleting a new project
func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectid") // getting id of project from the route
	fmt.Println("project deleted!", projectID)
	w.WriteHeader(http.StatusNoContent)
}

// CreateColumn is for creating a new Column
func (h *Handler) CreateColumn(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Column created!")
	fmt.Println(reqBody)
	h.App.CreateColumn(reqBody)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// ReadColumn is for reading a  Column
func (h *Handler) ReadColumn(w http.ResponseWriter, r *http.Request) {
	columnID := chi.URLParam(r, "columnid") // getting id of project from the route
	p, err := h.App.ReadColumn(columnID)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("this is ur column! ", p)
	w.WriteHeader(http.StatusOK)
}

// UpdateColumn is for updating a Column
func (h *Handler) UpdateColumn(w http.ResponseWriter, r *http.Request) {
	columnID := chi.URLParam(r, "columnid") // getting id of project from the route
	reqBody, _ := ioutil.ReadAll(r.Body)
	h.App.UpdateColumn(columnID, reqBody)
	fmt.Println("column updated!", columnID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteColumn is for deleting a Column
func (h *Handler) DeleteColumn(w http.ResponseWriter, r *http.Request) {
	columnID := chi.URLParam(r, "columnid") // getting id of project from the route
	fmt.Println("column deleted!", columnID)
	w.WriteHeader(http.StatusNoContent)
}

// CreateTask is for creating a new Task
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Task created!")
	fmt.Println(reqBody)
	h.App.CreateTask(reqBody)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// ReadTask is for reading a  Task
func (h *Handler) ReadTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskid") // getting id of Task from the route
	p, err := h.App.ReadTask(taskID)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("this is ur Task! ", p)
	w.WriteHeader(http.StatusOK)
}

// UpdateTask is for updating a Task
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskid") // getting id of Task from the route
	reqBody, _ := ioutil.ReadAll(r.Body)
	h.App.UpdateTask(taskID, reqBody)
	fmt.Println("Task updated!", taskID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteTask is for deleting a Task
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskid") // getting id of Comment from the route
	fmt.Println("Task deleted!", taskID)
	w.WriteHeader(http.StatusNoContent)
}

// CreateComment is for creating a new Comment
func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Comment created!")
	fmt.Println(reqBody)
	h.App.CreateComment(reqBody)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// ReadComment is for reading a  Comment
func (h *Handler) ReadComment(w http.ResponseWriter, r *http.Request) {
	commentID := chi.URLParam(r, "commentid") // getting id of Comment from the route
	p, err := h.App.ReadComment(commentID)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("this is ur comment! ", p)
	w.WriteHeader(http.StatusOK)
}

// UpdateComment is for updating a Comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	commentID := chi.URLParam(r, "commentid") // getting id of comment from the route
	reqBody, _ := ioutil.ReadAll(r.Body)
	h.App.UpdateComment(commentID, reqBody)
	fmt.Println("comment updated!", commentID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteComment is for deleting a Comment
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentID := chi.URLParam(r, "commentid") // getting id of comment from the route
	fmt.Println("comment deleted!", commentID)
	w.WriteHeader(http.StatusNoContent)
}
