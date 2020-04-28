package main

import (
	"a-go-project/db"
	"a-go-project/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var d *sql.DB = db.InitDB()

type Customers struct {
	Customers []models.Customer `json:"Customers"`
}

func InsertData() error {
	f, err := os.Open("input.json")
	if err != nil {
		log.Println("ERROR WHILE READING FILE ...")
		return err
	}
	jsonToByteArray, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("ERROR WHILE UNMARSHALLING JSON ...")
		return err
	}

	var cs Customers
	json.Unmarshal(jsonToByteArray, &cs)
	for i := 0; i < len(cs.Customers); i++ {
		newCustomer := models.Customer{}
		newCustomer.Name = cs.Customers[i].Name
		newCustomer.Address = cs.Customers[i].Address
		newCustomer.Age = cs.Customers[i].Age
		newCustomer.Email = cs.Customers[i].Email
		err := db.InsertCustomer(d, newCustomer)
		if err != nil {
			log.Println("ERROR INSERTING CUSTOMER ... ")
			return err
		}
	}
	return nil
}
func main() {
	//Clear table
	db.GiveMeDeath()
	//Insert data
	InsertData()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/add", add).Methods("POST")
	router.HandleFunc("/customers", ShowMeEveryone).Methods("GET")
	router.HandleFunc("/customer/{email}", GetOneCustomer).Methods("GET")
	router.HandleFunc("/removeCustomer/{email}", DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/updateCustomer/{email}", UpdateCustomer).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<h1>Test</h1>
	`)
}
func add(w http.ResponseWriter, r *http.Request) {
	var c models.Customer
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid JSON")
		log.Println(err)
		return
	}
	err = db.InsertCustomer(d, c)
	if err != nil {
		fmt.Fprintf(w, "error:", err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
}
func ShowMeEveryone(w http.ResponseWriter, r *http.Request) {
	var arr []models.Customer = db.GetAllCustomers(d)
	for _, p := range arr {
		fmt.Fprint(w, p)
	}
}
func GetOneCustomer(w http.ResponseWriter, r *http.Request) {
	customerEmail := mux.Vars(r)["email"]
	c, err := db.GetSpecificCustomer(d, customerEmail)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, c)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerEmail := mux.Vars(r)["email"]
	db.RemoveSpecificCustomer(customerEmail)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	customerEmail := mux.Vars(r)["email"]

	var c models.Customer
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid JSON")
		log.Println(err)
		return
	}

	if c.Email != customerEmail {
		fmt.Fprintf(w, "CUSTOMER DOES NOT EXIST ...")
		log.Println("CUSTOMER DOES NOT EXIST ...")
		return
	}

	rowsAffected, _ := db.UpdateCustomer(d, c)
	fmt.Fprintf(w, "CUSTOMER %s SUCCESSFULLY UPDATED ...", c.Email)
	log.Println(rowsAffected)
}
