package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	//logs
	infoLog := log.New(os.Stdout, "info: ", log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stdout, "error: ", log.Ltime|log.Lshortfile)

	//Initialize Database
	db, err := startDB()
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	//Dependencies
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}
	serve := &http.Server{
		ErrorLog: errorLog,
		Addr:     ":8080",
		Handler:  app.routes(),
	}
	//Start server
	infoLog.Println("STARTING SERVER ON PORT 8080: ...")
	err = serve.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}

}

//Database - Connect to DB
func startDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "web1:pass@tcp(localhost:3306)/videogames")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("DATABASE SUCCESSFULLY CONNECTED")
	return db, nil
}
