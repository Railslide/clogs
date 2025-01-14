package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	// "text/template"

	"github.com/railslide/clogs/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	sportRoutes, err := app.sportRoutes.Ongoing()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, sportRoute := range sportRoutes {
		fmt.Fprintf(w, "%+v\n", sportRoute)
	}
	// files := []string{
	// 	"./ui/html/base.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html",
	// 	"./ui/html/pages/home.tmpl.html",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
	//
	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
}

func (app *application) routeView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	sportRoute, err := app.sportRoutes.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", sportRoute)
}

func (app *application) routeCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to add a route..."))
}

func (app *application) routeCreatePost(w http.ResponseWriter, r *http.Request) {
	name := "The pink one in the corner"
	grade := "6A"
	routeType := "Top rope"
	sent := false
	archived := false
	gymID := 1
	id, err := app.sportRoutes.Insert(name, grade, routeType, gymID, sent, archived, "")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/route/view/%d", id), http.StatusSeeOther)
}

func (app *application) routeUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to update a route..."))
}
