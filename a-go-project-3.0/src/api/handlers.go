package api

import (
	"fmt"
	"net/http"

	"github.com/ksbeasle/GoLang/db/mysql"
)

/* Home -  This will make a call to the DB to get all the Games*/
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

/* GetGame - This will get one game based on the ID passed in else return errNoGameFound */
func GetGame(w http.ResponseWriter, r *http.Request) {
	// gorilla or regular?? --- id := mux.Vars(r)["id"]
	g, err := mysql.Get(1)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: %s ", err)
	}
	fmt.Fprintf(w, "", g)
}

/* AddGame - */
func AddGame(w http.ResponseWriter, r *http.Request) {

}

/* DeleteGame - */
func DeleteGame(w http.ResponseWriter, r *http.Request) {

}
