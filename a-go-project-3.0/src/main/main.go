package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ksbeasle/GoLang/api"
)

/* application struct - will hold our custom loggers and a db */
// type application struct {
// 	infoLog  *log.Logger
// 	errorLog *log.Logger
// 	db       *mysql.DBModel
// }

func main() {
	/* LOGS */
	infoLog := log.New(os.Stdout, "info: ", log.Lshortfile)
	errorLog := log.New(os.Stdout, "error: ", log.Lshortfile)

	/* Start a Connection to the Database - mysql */
	DB, err := startDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer DB.Close()
	// /* Create a new application struct */
	// app := &application{
	// 	infoLog:  infoLog,
	// 	errorLog: errorLog,
	// 	db:       &mysql.DBModel{DB: DB},
	// 	//gonna check something here later -- db:       DB,
	// }

	/* server struct */
	server := &http.Server{
		Addr:    ":8080",
		Handler: api.Routes(),
	}

	infoLog.Println("STARTING SERVER AT PORT ... ", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

/* startDB - Connect to the mysql database, return the db if successful else an error */
func startDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "web3:pass@tcp(localhost:3306)/videogames3")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("CONNCETION TO DATABASE SUCCESSFUL")
	return db, nil
}
