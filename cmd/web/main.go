package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

        // Static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
        mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /route/view/{id}", routeView)
	mux.HandleFunc("GET /route/create", routeCreate)
	mux.HandleFunc("POST /route/create", routeCreatePost)
	mux.HandleFunc("GET /route/edit/{id}", routeUpdate)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
