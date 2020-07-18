package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ksbeasle/GoLang/db/mysql"
)

/* Home -  */
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

/* GetGame - */
func GetGame(w http.ResponseWriter, r *http.Request) {
	g, err := mysql.Get(0)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "", g)
}

/* AddGame - */
func AddGame(w http.ResponseWriter, r *http.Request) {

}
