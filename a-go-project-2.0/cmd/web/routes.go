package main

import "net/http"

//Routes - REST
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.index)

	return mux
}
