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

//TODO: Custom error for query
//TODO: transaction maybe
func InitDB() (db *sql.DB) {
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	port = os.Getenv("PORT")
	dbConnectionString = fmt.Sprintf("%s:%s@tcp(localhost:%s)/test", user, password, port)

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func InsertCustomer(c models.Customer) {
	log.Println("INSERTING ...")

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("INSERT INTO test.customer (Name, Age, Email, Address) VALUES(?, ?, ?, ?)")
	if err != nil {
		return
	}

	_, err = stmt.Exec(c.Name, c.Age, c.Email, c.Address)
	if err != nil {
		return
	}
	defer stmt.Close()

}

func GetAllCustomers(db *sql.DB) []models.Customer {
	log.Println("GETTING ALL CUSTOMERS ...")

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err.Error())
	}

	var customersList []models.Customer
	rows, _ := db.Query("SELECT Name, Age, Email, Address FROM test.customer")

	var Name string
	var Age int
	var Email string
	var Address string
	for rows.Next() {
		err := rows.Scan(&Name, &Age, &Email, &Address)

		if err = rows.Err(); err != nil {
			//sql.Errnorows?
		}
		customersList = append(customersList, models.Customer{Name: Name, Age: Age, Email: Email, Address: Address})
	}
	err = rows.Close()
	if err = rows.Close(); err != nil {
		log.Println(err)
	}
	log.Println("ALL CUSTOMERS RETRIEVED")
	return customersList
}

func GetSpecificCustomer(db *sql.DB, email string) models.Customer {
	log.Println("GETTING DATA FROM SPECIFIC CUSTOMER ... ")
	var c models.Customer

	rows := db.QueryRow("SELECT Name, Age, Email, Address FROM test.customer WHERE Email=?", email)
	var Name string
	var Age int
	var Email string
	var Address string
	err := rows.Scan(&Name, &Age, &Email, &Address)
	//TODO: move this up to get all customers as well
	//TODO: add better errors
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No customer with that email address ...")
			panic(err.Error())
		} else {
			panic(err.Error())
		}
	}
	c = models.Customer{Name: Name, Age: Age, Email: Email, Address: Address}

	return c
}

func RemoveSpecificCustomer(email string) {
	log.Println("DELETING RECORD ... ")
	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("DELETE FROM test.customer WHERE Email = ?", email)
	if err != nil {
		log.Println("NO RECORD FOUND ... ")
		panic(err.Error())
	}

}

//Clear table(s)
func GiveMeDeath() string {
	log.Println("GOODBYE ... ")
	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Query("DROP TABLE customer")
	if err != nil {
		panic(err.Error())
	}

	return "... DB FLATLINE ..."
}
