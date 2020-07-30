package application

import (
	"database/sql"
	"log"

	"github.com/ksbeasle/GoLang/database"
)

type App struct {
	DBMODEL *database.GameDB
}

/*startDB - Connect to the mysql database, return the db if successful else an error */
func StartDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "web3:pass@tcp(localhost:3306)/videogames3?parseTime=true")
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
