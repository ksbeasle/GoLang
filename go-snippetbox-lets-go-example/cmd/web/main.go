package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"ksbeasle.net/snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

//application struct for app-wide dependencies
type application struct {
	InfoLog  *log.Logger
	errorLog *log.Logger
	snippets *mysql.SnippetModel
}

func main() {

	//get address from command line
	address := flag.String("addr", ":8080", "PORT")

	//DSN - parseTime is driver specific and helps convert sql time to Go time.Time
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	//Parse BEFORE using address
	flag.Parse()

	/* Custom loggers INFO/ERROR*/
	//using log.New(destination, string prefix, flags with additional info combined with | )

	//INFO
	InfoLog := log.New(os.Stdout, "info: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	//ERROR
	errorLog := log.New(os.Stdout, "error: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	//db
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	//intialize new instance of application
	//added snippetmodels so handlers can use it
	//dependencies
	app := &application{
		InfoLog:  InfoLog,
		errorLog: errorLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	//Server struct to use the new ERROR log that was created above
	serve := &http.Server{
		Addr:     *address,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	InfoLog.Printf("STARTING SERVER ON PORT %s ...", *address)
	//err := http.ListenAndServe(*address, mux)
	err = serve.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

/* Main() responsibility reduced to:
Parsing the runtime configuration settings for the application;
Establishing the dependencies for the handlers;
Running the HTTP server;
*/
