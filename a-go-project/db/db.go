package db

import (
	"a-go-project/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//TODO: Custom error for query maybe
func InitDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("PORT")
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(localhost:%s)/test", user, password, port)
	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err.Error())
	}
	log.Println(db.Ping())
}

func InsertCustomer(c models.Customer) models.Customer {
	fmt.Println(c.Name)
	stmt, err := db.Prepare("INSERT INTO test.user VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(c.Name, c.Age, c.Email, c.Address)
	if err != nil {
		panic(err.Error())
	}
	return c
}
