package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// generic internal server 500 error
func (app *applications) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *applications) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *applications) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
