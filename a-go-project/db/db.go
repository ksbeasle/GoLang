package db

import (
	"database/sql"
	"os"
)

func initDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("PORT")
	database := os.Getenv("GO_DB")
	db, err := sql.Open("mysql", "user:password@tcp(localhost:port)/database")
	if err != nil {
		panic(err.Error())
	}

}
