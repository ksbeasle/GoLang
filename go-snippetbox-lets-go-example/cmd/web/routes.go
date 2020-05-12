package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

//changed the return of the routes function to an http.Handler
//http.Handler instead of *http.ServeMux
func (app *application) routes() http.Handler {

	//Alice middleware package
	//The middleware here will be used on each request
	mw := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	//mux := http.NewServeMux()
	// mux.HandleFunc("/", app.home)
	// mux.HandleFunc("/snippet", app.showSnippet)
	// mux.HandleFunc("/snippet/create", app.createSnippet)

	//pat
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	//static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	//Old way -- Without Alice package
	//return app.recoverPanic(app.logRequest(secureHeaders(mux)))

	//With alice package
	return mw.Then(mux)

}
