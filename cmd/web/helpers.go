package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/m-golang/weather-app/internal/models"
)

// serverError handles internal server errors and logs the error stack.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}
// clientError handles client-side errors (e.g., 404, 400).
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)

}
// notFound handles 404 errors when a resource is not found.
func (app *application) notFound(w http.ResponseWriter) {
	data := app.newTempalteData(nil)

	// app.clientError(w, http.StatusNotFound)
	app.render(w, http.StatusNotFound, "404.html", data)
}
// newTempalteData prepares a new templateData structure with the current year.
func (app *application) newTempalteData(r *http.Request) *templateData {
	return &templateData{
		Weather:     &models.Weather{},
		CurrentYear: time.Now().Year(),
	}
}
// render renders a template with the given status and data.
func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}
