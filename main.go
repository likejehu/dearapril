package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/likejehu/dearapril/handlers"
)

func main() {
	handler := handlers.Handler{}
	r := chi.NewRouter()
	//middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//routing
	r.Get("/projects", handler.GetAllProjects)
	r.Post("/{projectid}", handler.Hello)
	r.Get("/{projectid}", handler.Hello)
	r.Patch("/{projectid}", handler.Hello)
	r.Delete("/{projectid}", handler.Hello)

	r.Post("/{projectid}/{columnid}", handler.Hello)
	r.Get("/{projectid}/{columnid}", handler.Hello)
	r.Patch("/{projectid}/{columnid}", handler.Hello)
	r.Delete("/{projectid}/{columnid}", handler.Hello)

	r.Post("/{projectid}/{columnid}/{taskid}", handler.Hello)
	r.Get("/{projectid}/{columnid}/{taskid}", handler.Hello)
	r.Patch("/{projectid}/{columnid}/{taskid}", handler.Hello)
	r.Delete("/{projectid}/{columnid}/{taskid}", handler.Hello)

	r.Post("/{projectid}/{columnid}/{taskid}/{commentid}", handler.Hello)
	r.Get("/{projectid}/{columnid}/{taskid}/{commentid}", handler.Hello)
	r.Patch("/{projectid}/{columnid}/{taskid}/{commentid}", handler.Hello)
	r.Delete("/{projectid}/{columnid}/{taskid}/{commentid}", handler.Hello)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", r))
}
