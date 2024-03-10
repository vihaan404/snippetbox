package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type snippet struct {
	str string
}

func (app *applications) handlerHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// log.Print(err.Error())
		app.errorLog.Printf(err.Error())
		http.Error(w, "internal server error", 500)
		return

	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {

		app.errorLog.Printf(err.Error())
		http.Error(w, "internal server error", 500)

	}
	w.Write([]byte("Hello from the snippetbox"))
}

// display snippet
func (app *applications) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// creating snippet

func (app *applications) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("This medhod %s is not allowed"))
		http.Error(w, "this is method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create new snippet"))
}
