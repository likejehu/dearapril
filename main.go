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
	r.Get("/hello", handler.Hello)
	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", r))
}
