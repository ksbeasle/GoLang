package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ksbeasle/GoLang/database"
)

/* Home -  This will make a call to the DB to get all the Games*/
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

/* GetGame - This will get one game based on the ID passed in else return errNoGameFound */
func GetGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "error: ", err)
	}
	g, err := database.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: %s ", err)
	}
	fmt.Fprintf(w, "GAMe: ", g)
}

/* AddGame - */
func AddGame(w http.ResponseWriter, r *http.Request) {

}

/* DeleteGame - */
func DeleteGame(w http.ResponseWriter, r *http.Request) {

}
