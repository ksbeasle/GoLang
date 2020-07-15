package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ksbeasle/GoLang/db/mysql"
)

/* Application struct - will hold our custom loggers and a db */
type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *mysql.DBModel
}

func main() {
	/* LOGS */
	infoLog := log.New(os.Stdout, "info: ", log.Lshortfile)
	errorLog := log.New(os.Stdout, "error: ", log.Lshortfile)

	/* Start a Connection to the Database - mysql */
	DB, err := startDB()
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

	err := db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println(db.Ping())
	log.Println("CONNCETION TO DATABASE SUCCESSFUL")
	return db, nil
}
