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

	//We create a new middleware for the sessions - nosurf prevent attacks - app authenticate to pass context of authenicated user if present
	dmw := alice.New(app.session.Enable, noSurf, app.authenticate)

	//mux := http.NewServeMux()
	// mux.HandleFunc("/", app.home)
	// mux.HandleFunc("/snippet", app.showSnippet)
	// mux.HandleFunc("/snippet/create", app.createSnippet)

	//pat - router
	mux := pat.New()
	mux.Get("/", dmw.ThenFunc(app.home))
	mux.Get("/snippet/create", dmw.Append(app.requireAuthentication).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dmw.Append(app.requireAuthentication).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dmw.ThenFunc(app.showSnippet))

	mux.Get("/user/signup", dmw.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dmw.ThenFunc(app.signupUser))
	mux.Get("/user/login", dmw.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dmw.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dmw.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	mux.Get("/ping", http.HandlerFunc(ping))

	//static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	//Old way -- Without Alice package
	//return app.recoverPanic(app.logRequest(secureHeaders(mux)))

	//With alice package
	return mw.Then(mux)

}
