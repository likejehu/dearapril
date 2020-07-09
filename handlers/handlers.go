package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/likejehu/dearapril/models"
	"github.com/likejehu/dearapril/validation"
)

// AppController is interface for main logic
type AppController interface {
	CreateProject(p *models.Project) (projid int, err error)
	UpdateProject(id int, p *models.Project) (err error)
	ReadProject(id int) (p *models.Project, err error)
	DeleteProject(id int) (err error)
	ReadProjects() (p []*models.Project, err error)

	CreateColumn(col *models.Column, projid int) (colid int, err error)
	ReadColumn(id int) (col *models.Column, err error)
	UpdateColumn(id int, col *models.Column) (colid int, err error)
	DeleteColumn(id int) (err error)
	ReadColumns() (col []*models.Column, err error)
	MoveColumn(id int, direction string)

	CreateTask(t *models.Task, colid int) (taskid int, err error)
	ReadTask(id int) (t *models.Task, err error)
	UpdateTask(id int, t *models.Task) (err error)
	DeleteTask(id int) (err error)
	ReadTasks() (tasks []*models.Task, err error)
	MoveTaskToColumn(colid, taskid int) (err error)
	MoveTaskUpDown(taskid int, direction string) (err error)

	CreateComment(com *models.Comment, taskid int) (comid int, err error)
	ReadComment(id int) (com *models.Comment, err error)
	UpdateComment(id int, com *models.Comment) (err error)
	DeleteComment(id int) (err error)
	ReadComments() (com []*models.Comment, err error)
}

// Handler is struct for handlers
type Handler struct {
	App AppController
}

// GetAllProjects gets all the projets
func (h *Handler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.App.ReadProjects()
	if err != nil {
		log.Print(err)
		return
	}
	js, err := json.Marshal(projects)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	return

}

// CreateProject is for creating a new project
func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	p := new(models.Project)
	json.Unmarshal(reqBody, &p)
	validation.ValidateStruct(p)
	h.App.CreateProject(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// ReadProject is for reading a new project
func (h *Handler) ReadProject(w http.ResponseWriter, r *http.Request) {
	pID := chi.URLParam(r, "projectid") // getting id of project from the route
	projectID, _ := strconv.Atoi(pID)
	p, err := h.App.ReadProject(projectID)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(p)
	if err != nil {
		log.Print(err)
		return
	}
	w.Write(js)
	w.WriteHeader(http.StatusOK)
	return
}

// UpdateProject is for updating a new project
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	pID := chi.URLParam(r, "projectid") // getting id of project from the route
	reqBody, _ := ioutil.ReadAll(r.Body)
	p := new(models.Project)
	json.Unmarshal(reqBody, &p)
	validation.ValidateStruct(p)
	projectID, _ := strconv.Atoi(pID)
	h.App.UpdateProject(projectID, p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// DeleteProject is for deleting a new project
func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	pID := chi.URLParam(r, "projectid") // getting id of project from the route
	projectID, _ := strconv.Atoi(pID)
	h.App.DeleteProject(projectID)
	w.WriteHeader(http.StatusNoContent)
	return
}

// GetAllColumns gets all the columns
func (h *Handler) GetAllColumns(w http.ResponseWriter, r *http.Request) {
	columns, err := h.App.ReadColumns()
	js, err := json.Marshal(columns)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	return

}

// CreateColumn is for creating a new Column
func (h *Handler) CreateColumn(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	pID := chi.URLParam(r, "projectid") // getting id of column from the route
	c := new(models.Column)
	json.Unmarshal(reqBody, &c)
	validation.ValidateStruct(c)
	projectID, _ := strconv.Atoi(pID)
	h.App.CreateColumn(c, projectID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// ReadColumn is for reading a  Column
func (h *Handler) ReadColumn(w http.ResponseWriter, r *http.Request) {
	cID := chi.URLParam(r, "columnid") // getting id of column from the route
	columnID, _ := strconv.Atoi(cID)
	c, err := h.App.ReadColumn(columnID)
	if err != nil {
		log.Println(err)
		return
	}
	js, err := json.Marshal(c)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	return
}

// UpdateColumn is for updating a Column
func (h *Handler) UpdateColumn(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	cID := chi.URLParam(r, "columnid") // getting id of column from the route
	columnID, _ := strconv.Atoi(cID)
	c := new(models.Column)
	json.Unmarshal(reqBody, &c)
	validation.ValidateStruct(c)
	h.App.UpdateColumn(columnID, c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// DeleteColumn is for deleting a Column
func (h *Handler) DeleteColumn(w http.ResponseWriter, r *http.Request) {
	cID := chi.URLParam(r, "columnid") // getting id of column from the route
	columnID, _ := strconv.Atoi(cID)
	h.App.DeleteColumn(columnID)
	taskID := 1
	h.App.MoveTaskToColumn(columnID, taskID)
	w.WriteHeader(http.StatusNoContent)
	return
}

//MoveColumn is for moving given column across project
func (h *Handler) MoveColumn(w http.ResponseWriter, r *http.Request) {
	cID := chi.URLParam(r, "columnid")        // getting id of column from the route
	direction := chi.URLParam(r, "direction") // getting direction  from the route
	columnID, _ := strconv.Atoi(cID)
	h.App.MoveColumn(columnID, direction)
	w.WriteHeader(http.StatusOK)
	return
}

// GetAllTasks gets all the Tasks
func (h *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.App.ReadTasks()
	js, err := json.Marshal(tasks)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	return

}

// CreateTask is for creating a new Task
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	cID := chi.URLParam(r, "columnid") // getting id of column from the route
	t := new(models.Task)
	json.Unmarshal(reqBody, &t)
	validation.ValidateStruct(t)
	columnID, _ := strconv.Atoi(cID)
	h.App.CreateTask(t, columnID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// ReadTask is for reading a  Task
func (h *Handler) ReadTask(w http.ResponseWriter, r *http.Request) {
	tID := chi.URLParam(r, "taskid") // getting id of Task from the route
	taskID, _ := strconv.Atoi(tID)
	t, err := h.App.ReadTask(taskID)
	if err != nil {
		log.Println(err)
		return
	}
	js, err := json.Marshal(t)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	return
}

// UpdateTask is for updating a Task
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	tID := chi.URLParam(r, "taskid") // getting id of Task from the route
	taskID, _ := strconv.Atoi(tID)
	reqBody, _ := ioutil.ReadAll(r.Body)
	t := new(models.Task)
	json.Unmarshal(reqBody, &t)
	validation.ValidateStruct(t)
	h.App.UpdateTask(taskID, t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// DeleteTask is for deleting a Task
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	tID := chi.URLParam(r, "taskid") // getting id of Task from the route
	taskID, _ := strconv.Atoi(tID)
	h.App.DeleteTask(taskID)
	w.WriteHeader(http.StatusNoContent)
	return
}

// MoveTaskToColumn is for movin a Task across columns
func (h *Handler) MoveTaskToColumn(w http.ResponseWriter, r *http.Request) {
	tID := chi.URLParam(r, "taskid") // getting id of Task from the route
	taskID, _ := strconv.Atoi(tID)
	cID := chi.URLParam(r, "newcolumnid") // getting id new column from the route
	columnID, _ := strconv.Atoi(cID)
	h.App.MoveTaskToColumn(taskID, columnID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// MoveTask is for movin a Task up and down
func (h *Handler) MoveTask(w http.ResponseWriter, r *http.Request) {
	tID := chi.URLParam(r, "taskid") // getting id of Task from the route
	taskID, _ := strconv.Atoi(tID)
	direction := chi.URLParam(r, "direction") // getting direction  from the route
	h.App.MoveTaskUpDown(taskID, direction)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// GetAllComments gets all the Comments
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.App.ReadComments()
	js, err := json.Marshal(comments)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	return

}

// CreateComment is for creating a new Comment
func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	comID := chi.URLParam(r, "commentid") // getting id of Comment from the route
	commentID, _ := strconv.Atoi(comID)
	com := new(models.Comment)
	json.Unmarshal(reqBody, &com)
	validation.ValidateStruct(com)
	h.App.CreateComment(com, commentID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// ReadComment is for reading a  Comment
func (h *Handler) ReadComment(w http.ResponseWriter, r *http.Request) {
	comID := chi.URLParam(r, "commentid") // getting id of Comment from the route
	commentID, _ := strconv.Atoi(comID)
	com, err := h.App.ReadComment(commentID)
	if err != nil {
		log.Println(err)
		return
	}
	js, err := json.Marshal(com)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	return
}

// UpdateComment is for updating a Comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comID := chi.URLParam(r, "commentid") // getting id of Comment from the route
	commentID, _ := strconv.Atoi(comID)
	reqBody, _ := ioutil.ReadAll(r.Body)
	com := new(models.Comment)
	json.Unmarshal(reqBody, &com)
	validation.ValidateStruct(com)
	h.App.UpdateComment(commentID, com)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return

}

// DeleteComment is for deleting a Comment
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	comID := chi.URLParam(r, "commentid") // getting id of comment from the route
	commentID, _ := strconv.Atoi(comID)
	h.App.DeleteComment(commentID)
	w.WriteHeader(http.StatusNoContent)
	return
}
