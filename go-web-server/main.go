package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var count int

var mutex = &sync.Mutex{}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index")
}

func increment(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	count++
	fmt.Fprintf(w, strconv.Itoa(count))
	mutex.Unlock()
}

func main() {

	//index
	//http.HandleFunc("/", index)

	//increment
	http.HandleFunc("/increment", increment)
	http.Handle("/", http.FileServer(http.Dir("./static-html-files")))
	//Serve HTML file
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	//fmt.Fprintf(w, "url - > %v", r.URL.Path[1:])
	// 	http.ServeFile(w, r, r.URL.Path[1:])
	// })
	http.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Endpoint - test1")
		fmt.Fprintf(w, "This is test1")
	})

	log.Fatal(http.ListenAndServe(":8888", nil))

	//HTTPS
	//GETTING AN ERROR HERE -- 2020/03/25 16:38:42 http: TLS handshake error from [::1]:57193: EOF
	//log.Fatal(http.ListenAndServeTLS(":666", "server.crt", "serverkey", nil))
}
