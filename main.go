package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/likejehu/dearapril/core"
	"github.com/likejehu/dearapril/db"
	"github.com/likejehu/dearapril/handlers"
)

func main() {

	conn := db.PostgreStore.DB
	defer conn.Close()

	handler := handlers.Handler{
		App: core.AppController,
	}

	r := chi.NewRouter()
	//middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//routing
	r.Get("/projects", handler.GetAllProjects)
	r.Post("/createproject", handler.CreateProject)
	r.Get("/projects/{projectid}", handler.ReadProject)
	r.Patch("/projects/{projectid}", handler.UpdateProject)
	r.Delete("/projects/{projectid}", handler.DeleteProject)

	r.Get("/projects/{projectid}/columns", handler.GetAllColumns)
	r.Post("/projects/{projectid}/createcolumn", handler.CreateColumn)
	r.Get("/projects/{projectid}/columns/{columnid}", handler.ReadColumn)
	r.Patch("/projects/{projectid}/columns/{columnid}", handler.UpdateColumn)
	r.Delete("/projects/{projectid}/columns/{columnid}/{leftid}", handler.DeleteColumn)
	r.Patch("/projects/{projectid}/columns/{columnid}/move/{nextid}", handler.MoveColumn)

	r.Get("/projects/{projectid}/columns/{columnid}/tasks", handler.GetAllTasks)
	r.Post("/projects/{projectid}/columns/{columnid}/createtask", handler.CreateTask)
	r.Get("/projects/{projectid}/columns/{columnid}/tasks/{taskid}", handler.ReadTask)
	r.Patch("/projects/{projectid}/columns/{columnid}/tasks/{taskid}", handler.UpdateTask)
	r.Delete("/projects/{projectid}/columns/{columnid}/tasks/{taskid}", handler.DeleteTask)
	r.Patch("/projects/{projectid}/columns/{columnid}/tasks/{taskid}/movetocolumn/{newcolumnid}", handler.MoveTaskToColumn)
	r.Patch("/projects/{projectid}/columns/{columnid}/tasks/{taskid}/move/{nextid}", handler.MoveTask)

	r.Get("/projects/{projectid}/columns/{columnid}/tasks/{taskid}/comments", handler.GetAllComments)
	r.Post("/projects/{projectid}/columns/{columnid}/tasks/{taskid}/createcomment", handler.CreateComment)
	r.Get("/projects/{projectid}/columns/{columnid}/tasks/{taskid}/comments/{commentid}", handler.ReadComment)
	r.Patch("/projects/{projectid}/columns/{columnid}/tasks/{taskid}/comments/{commentid}", handler.UpdateComment)
	r.Delete("/projects/{projectid}/columns/{columnid}/tasks/{taskid}/comments/{commentid}", handler.DeleteComment)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", r))
}
