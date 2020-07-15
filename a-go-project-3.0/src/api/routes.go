package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

/* Routes - Uses gorilla mux */
func Routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
}
