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

	//We create a new middleware for the sessions
	dmw := alice.New(app.session.Enable)

	//mux := http.NewServeMux()
	// mux.HandleFunc("/", app.home)
	// mux.HandleFunc("/snippet", app.showSnippet)
	// mux.HandleFunc("/snippet/create", app.createSnippet)

	//pat
	mux := pat.New()
	mux.Get("/", dmw.ThenFunc(app.home))
	mux.Get("/snippet/create", dmw.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dmw.ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dmw.ThenFunc(app.showSnippet))

	mux.Get("/user/signup", dmw.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dmw.ThenFunc(app.signupUser))
	mux.Get("/user/login", dmw.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dmw.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dmw.ThenFunc(app.logoutUser))

	//static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	//Old way -- Without Alice package
	//return app.recoverPanic(app.logRequest(secureHeaders(mux)))

	//With alice package
	return mw.Then(mux)

}
