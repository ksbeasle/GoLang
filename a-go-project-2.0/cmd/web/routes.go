package main

import "net/http"

//Routes - REST
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.index)        //HOME
	mux.HandleFunc("/game/add", app.add)  //ADD a new game
	mux.HandleFunc("/game/", app.getGame) //Get specific game

	return mux
}
