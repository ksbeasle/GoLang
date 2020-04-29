package db

import (
	"a-go-project/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB = InitDB()
var user string
var password string
var port string
var dbConnectionString string

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

func InsertCustomer(db *sql.DB, c models.Customer) error {
	log.Println("INSERTING ...")

	stmt, err := db.Prepare("INSERT INTO test.customer (Name, Age, Email, Address) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(c.Name, c.Age, c.Email, c.Address)
	if err != nil {
		return err
	}
	defer stmt.Close()
	log.Printf("CUSTOMER %s SUCCESSFULLY ADDED ...", c.Email)
	return nil

}

func UpdateCustomer(db *sql.DB, c models.Customer) (int64, error) {
	log.Println("ATTEMPTING TO UPDATE CUSTOMER -- ", c.Email)
	_, err := GetSpecificCustomer(db, c.Email)
	if err != nil {
		log.Println("CUSTOMER DOES NOT EXIST ...", err)
	}

	result, err := db.Exec("UPDATE test.customer SET Name = ?, Email = ?,Age = ?, Address = ? WHERE Email = ?", c.Name, c.Email, c.Age, c.Address, c.Email)
	return result.RowsAffected()
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

		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("No customer with that email address ...")
				log.Println("Error No Rows")
			} else {
				panic(err.Error())
			}
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

func GetSpecificCustomer(db *sql.DB, email string) (models.Customer, error) {
	log.Println("GETTING DATA FROM SPECIFIC CUSTOMER ... ")
	var c models.Customer

	rows := db.QueryRow("SELECT Name, Age, Email, Address FROM test.customer WHERE Email=?", email)
	var Name string
	var Age int
	var Email string
	var Address string
	err := rows.Scan(&Name, &Age, &Email, &Address)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No customer with that email address ...")
			log.Println("Error No Rows")
			return c, err
		} else {
			panic(err.Error())
		}
	}
	c = models.Customer{Name: Name, Age: Age, Email: Email, Address: Address}
	log.Printf("RECORD FOR %s SUCCESSFULLY RETRIEVED ...", c.Email)
	return c, nil
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
	log.Printf("RECORD FOR %s SUCCESSFULLY DELETED ...", email)

}

//Clear table
func GiveMeDeath() {
	log.Println("CLEARING TABLE ... ")
	_, err := db.Query("TRUNCATE TABLE test.customer")
	if err != nil {
		panic(err.Error())
	}
	log.Println("TABLE CLEARED SUCCESSFULLY ...")
}
