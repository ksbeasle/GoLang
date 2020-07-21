package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

/* Routes - Uses gorilla mux */
func Routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/game/{id}", GetGame).Methods("GET")
	r.HandleFunc("/game/add", AddGame).Methods("POST")
	r.HandleFunc("/game/delete/{id}", DeleteGame).Methods("DELETE")

	return r
}
