package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/vihaan404/snippetbox/internal/models"
)

type snippet struct {
	str string
}

func (app *applications) home(w http.ResponseWriter, r *http.Request) {
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
		// app.errorLog.Printf(err.Error())
		// http.Error(w, "internal server error", 500)
		app.serverError(w, err)
		return

	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {

		app.errorLog.Printf(err.Error())

		app.serverError(w, err)
	}

	w.Write([]byte("Hello from the snippetbox"))
}

// display snippet
func (app *applications) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecords) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return

	}

	fmt.Fprintf(w, "%+v", snippet)
}

// creating snippet

func (app *applications) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("This medhod %s is not allowed"))
		// http.Error(w, "this is method is not allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "O snail"
	expires := 7
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
