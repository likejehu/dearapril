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

	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
	r.Get("/{projectid}", handler.ReadProject)
	r.Patch("/{projectid}", handler.UpdateProject)
	r.Delete("/{projectid}", handler.DeleteProject)

	r.Get("/{projectid}/columns", handler.GetAllColumns)
	r.Post("/{projectid}/createcolumn", handler.CreateColumn)
	r.Get("/{projectid}/{columnid}", handler.ReadColumn)
	r.Patch("/{projectid}/{columnid}", handler.UpdateColumn)
	r.Delete("/{projectid}/{columnid}/{leftid}", handler.DeleteColumn)
	r.Patch("/{projectid}/{columnid}/move/{nextid}", handler.MoveColumn)

	r.Get("/{projectid}/{columnid}/tasks", handler.GetAllTasks)
	r.Post("/{projectid}/{columnid}/createtask", handler.CreateTask)
	r.Get("/{projectid}/{columnid}/{taskid}", handler.ReadTask)
	r.Patch("/{projectid}/{columnid}/{taskid}", handler.UpdateTask)
	r.Delete("/{projectid}/{columnid}/{taskid}", handler.DeleteTask)
	r.Patch("/{projectid}/{columnid}/{taskid}/movetocolumn/{newcolumnid}", handler.MoveTaskToColumn)
	r.Patch("/{projectid}/{columnid}/{taskid}/move/{nextid}", handler.MoveTask)

	r.Get("//{projectid}/{columnid}/{taskid}/comments", handler.GetAllComments)
	r.Post("/{projectid}/{columnid}/{taskid}/createcomment", handler.CreateComment)
	r.Get("/{projectid}/{columnid}/{taskid}/{commentid}", handler.ReadComment)
	r.Patch("/{projectid}/{columnid}/{taskid}/{commentid}", handler.UpdateComment)
	r.Delete("/{projectid}/{columnid}/{taskid}/{commentid}", handler.DeleteComment)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", r))
}
