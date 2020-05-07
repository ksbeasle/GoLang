package main

import "net/http"

//HOME
func (app *application) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

}

//Add a game
func (app *application) add(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("ADDING GAME ... ")

	//Only allow POST method
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)

	}
}
