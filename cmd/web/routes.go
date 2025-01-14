package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	// Static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /route/view/{id}", app.routeView)
	mux.HandleFunc("GET /route/create", app.routeCreate)
	mux.HandleFunc("POST /route/create", app.routeCreatePost)
	mux.HandleFunc("GET /route/edit/{id}", app.routeUpdate)

        return mux

}
