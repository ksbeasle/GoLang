package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"ksbeasle.net/snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

//application struct for app-wide dependencies
type application struct {
	InfoLog       *log.Logger
	errorLog      *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
	users         *mysql.UserModel
}

func main() {

	//get address from command line
	address := flag.String("addr", ":8080", "PORT")

	//DSN - parseTime is driver specific and helps convert sql time to Go time.Time
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	//Secret for the new session
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")

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

	//intialize new template cache
	// Initialize a new template cache...
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	//create new session and set the life span to 12 hours
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	//intialize new instance of application
	//added snippetmodels so handlers can use it
	//dependencies
	app := &application{
		InfoLog:       InfoLog,
		errorLog:      errorLog,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
		users:         &mysql.UserModel{DB: db},
	}
	//TLS config - cert/key stuff
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	//Server struct to use the new ERROR log that was created above
	serve := &http.Server{
		Addr:         *address,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	InfoLog.Printf(" ------------------------- STARTING SERVER ON PORT %s ... -------------------------", *address)
	//err := http.ListenAndServe(*address, mux)
	//err = serve.ListenAndServe()
	err = serve.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
