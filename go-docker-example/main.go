package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test")
	})

	log.Fatal(http.ListenAndServe(":7070", nil))
}
