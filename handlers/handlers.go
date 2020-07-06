package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/likejehu/dearapril/models"
)

// AppController is interface for main logic
type AppController interface {
	SayHello()
	CreateProject(p *models.Project) (projid int, err error)
	UpdateProject(id int, p *models.Project) (err error)
	ReadProject(id int) (p *models.Project, err error)
	DeleteProject(id int) (err error)

	CreateColumn(col *models.Column, projid int) (colid int, err error)
	ReadColumn(id int) (col *models.Column, err error)
	UpdateColumn(id int, col *models.Column) (colid int, err error)
	DeleteColumn(id int) (err error)

	CreateTask(t *models.Task, colid int) (taskid int, err error)
	ReadTask(id int) (t *models.Task, err error)
	UpdateTask(id int, t *models.Task)
	DeleteTask(id int) (err error)
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
	p := new(models.Project)
	json.Unmarshal(reqBody, &p)
	h.App.CreateProject(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// ReadProject is for reading a new project
func (h *Handler) ReadProject(w http.ResponseWriter, r *http.Request) (p *models.Project, err error) {
	pID := chi.URLParam(r, "projectid") // getting id of project from the route
	projectID, _ := strconv.Atoi(pID)
	p, err = h.App.ReadProject(projectID)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return p, err
}

// UpdateProject is for updating a new project
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	pID := chi.URLParam(r, "projectid") // getting id of project from the route
	reqBody, _ := ioutil.ReadAll(r.Body)
	p := new(models.Project)
	json.Unmarshal(reqBody, &p)
	projectID, _ := strconv.Atoi(pID)
	h.App.UpdateProject(projectID, p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteProject is for deleting a new project
func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	pID := chi.URLParam(r, "projectid") // getting id of project from the route
	projectID, _ := strconv.Atoi(pID)
	h.App.DeleteProject(projectID)
	w.WriteHeader(http.StatusNoContent)
}

// CreateColumn is for creating a new Column
func (h *Handler) CreateColumn(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	pID := chi.URLParam(r, "projectid") // getting id of column from the route
	c := new(models.Column)
	json.Unmarshal(reqBody, &c)
	projectID, _ := strconv.Atoi(pID)
	h.App.CreateColumn(c, projectID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// ReadColumn is for reading a  Column
func (h *Handler) ReadColumn(w http.ResponseWriter, r *http.Request) (c *models.Column, err error) {
	cID := chi.URLParam(r, "columnid") // getting id of column from the route
	columnID, _ := strconv.Atoi(cID)
	c, err = h.App.ReadColumn(columnID)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	return c, err
}

// UpdateColumn is for updating a Column
func (h *Handler) UpdateColumn(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	cID := chi.URLParam(r, "columnid") // getting id of column from the route
	columnID, _ := strconv.Atoi(cID)
	c := new(models.Column)
	json.Unmarshal(reqBody, &c)

	h.App.UpdateColumn(columnID, c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteColumn is for deleting a Column
func (h *Handler) DeleteColumn(w http.ResponseWriter, r *http.Request) {
	cID := chi.URLParam(r, "columnid") // getting id of column from the route
	columnID, _ := strconv.Atoi(cID)
	h.App.DeleteColumn(columnID)
	w.WriteHeader(http.StatusNoContent)
}

// CreateTask is for creating a new Task
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	cID := chi.URLParam(r, "columnid") // getting id of column from the route
	t := new(models.Task)
	json.Unmarshal(reqBody, &t)
	columnID, _ := strconv.Atoi(cID)
	h.App.CreateTask(t, columnID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// ReadTask is for reading a  Task
func (h *Handler) ReadTask(w http.ResponseWriter, r *http.Request) (t *models.Task, err error) {
	tID := chi.URLParam(r, "taskid") // getting id of Task from the route
	taskID, _ := strconv.Atoi(tID)
	t, err = h.App.ReadTask(taskID)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	return t, err
}

// UpdateTask is for updating a Task
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	tID := chi.URLParam(r, "taskid") // getting id of Task from the route
	taskID, _ := strconv.Atoi(tID)
	reqBody, _ := ioutil.ReadAll(r.Body)
	t := new(models.Task)

	json.Unmarshal(reqBody, &t)
	h.App.UpdateTask(taskID, t)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteTask is for deleting a Task
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	tID := chi.URLParam(r, "taskid") // getting id of Task from the route
	taskID, _ := strconv.Atoi(tID)
	h.App.DeleteTask(taskID)
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
