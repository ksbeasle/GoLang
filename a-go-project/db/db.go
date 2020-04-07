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
var user string
var password string
var port string
var dbConnectionString string

//TODO: Custom error for query maybe

func InitDB() {
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	port = os.Getenv("PORT")
	dbConnectionString = fmt.Sprintf("%s:%s@tcp(localhost:%s)/test", user, password, port)
}

func InsertCustomer(c models.Customer) {
	log.Println("INSERTING ...")

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("INSERT INTO test.customer (Name, Age, Email, Address) VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(c.Name, c.Age, c.Email, c.Address)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	defer db.Close()

}

func GetAllCustomers() []models.Customer {
	log.Println("GETTING ALL CUSTOMERS ...")

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err.Error())
	}

	var customersList []models.Customer
	rows, err := db.Query("SELECT * FROM test.customer")
	if err != nil {
		panic(err.Error())
	}
	var Name string
	var Age int
	var Email string
	var Address string
	//TODO: ERROR BELOW
	//panic: sql: expected 5 destination arguments in Scan, not 4
	for rows.Next() {
		err := rows.Scan(&Name, &Age, &Email, &Address)
		if err != nil {
			panic(err.Error())
		}
		customersList = append(customersList, models.Customer{Name: Name, Age: Age, Email: Email, Address: Address})
	}
	defer db.Close()
	return customersList
}
