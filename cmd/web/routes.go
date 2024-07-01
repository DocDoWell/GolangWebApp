package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux:= chi.NewRouter()

	//keep app up and running
	mux.Use(middleware.Recoverer)
	//timeout after 1 min
	mux.Use(middleware.Timeout(60*time.Second))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	//display our test pages
	mux.Get("/test-patterns", app.TestPatterns)

	//factory routes
	mux.Get("/api/dog-from-factory", app.CreateDogFromFactory)
	mux.Get("/api/cat-from-factory", app.CreateCatFromFactory)

	mux.Get("/api/dog-from-abstract-factory", app.createDogFromAbstractFactory)
	mux.Get("/api/cat-from-abstract-factory", app.createCatFromAbstractFactory)


	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)

	return mux
}
