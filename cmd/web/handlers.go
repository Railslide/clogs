package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
	}
}

func routeView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific route with id %d...", id)
	w.Write([]byte(msg))
}

func routeCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to add a route..."))
}

func routeCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new route..."))
}

func routeUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to update a route..."))
}
