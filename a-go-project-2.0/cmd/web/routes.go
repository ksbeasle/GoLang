package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

//Routes - REST
func (app *application) routes() http.Handler {
	//Using Pat router
	//since Pat does not allow us to register handlerFunc directly
	//We have to do it ourselves
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.index))           //HOME - GET all
	mux.Post("/game/add", http.HandlerFunc(app.add))    // ADD a new game
	mux.Get("/game/:id", http.HandlerFunc(app.getGame)) //Get specific game

	return app.LogIncomingRequests(mux)
}
