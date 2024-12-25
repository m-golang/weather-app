package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes configures the application's HTTP routes.
func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Set up the 404 route handler
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	// Static file server for CSS, JavaScript, and other assets
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	// Routes for the home page and weather view
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/weather/:city", app.weatherView)

	// Apply middlewares
	return app.recoverPanic(app.logRequest(secureHeaders(router)))
}
