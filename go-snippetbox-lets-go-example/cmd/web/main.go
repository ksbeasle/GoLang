package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	//get address from command line
	address := flag.String("addr", ":8080", "PORT")

	//Parse BEFORE using address
	flag.Parse()

	/* Custom loggers INFO/ERROR*/
	//using log.New(destination, string prefix, flags with additional info combined with | )

	//INFO
	InfoLog := log.New(os.Stdout, "info: ", log.Ldate|log.Ltime|log.LUTC|log.Lmicroseconds|log.Lshortfile)

	//ERROR
	errorLog := log.New(os.Stdout, "error: ", log.Ldate|log.Ltime|log.LUTC|log.Lmicroseconds|log.Lshortfile)

	/* PLACE TO KEEP OUR LOGS */
	f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer f.Close()
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Server struct to use the new ERROR log that was created above
	serve := &http.Server{
		Addr:     *address,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	InfoLog.Printf("STARTING SERVER ON PORT %s ...", *address)
	//err := http.ListenAndServe(*address, mux)
	err = serve.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
