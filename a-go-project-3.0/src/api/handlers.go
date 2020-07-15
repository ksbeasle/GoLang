package api

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write("")
}
