package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) routeView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific route with id %d...", id)
	w.Write([]byte(msg))
}

func (app *application) routeCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to add a route..."))
}

func (app *application) routeCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new route..."))
}

func (app *application) routeUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to update a route..."))
}
