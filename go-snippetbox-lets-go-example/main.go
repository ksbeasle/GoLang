package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)
//http://localhost:8080/snippet?id=123


func Home(w http.ResponseWriter, r *http.Request){
	//this will prevent the slash catching all requests
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("Hello World"))
}

func showSnippet(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		//http.Error(w, "Invalid ID", http.StatusNotFound)
		http.NotFound(w,r)
		return
	}
	fmt.Fprintf(w, "Id = %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request){
	//Allow only POST method to be used per HTTP good practices
	if r.Method != http.MethodPost {
		//Let user know what Methods are accepted at the endpoint
		w.Header().Set("Allow", http.MethodPost)
		//405
		http.Error(w, "Method Now allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a snippet"))
}


func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("STARTING SERVER ON PORT 8080 ...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil{
		log.Fatal(err)
	}
}