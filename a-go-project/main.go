package main

import (
	"a-go-project/db"
	"a-go-project/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Customers struct {
	Customers []models.Customer `json:"Customers"`
}

func main() {

	f, err := os.Open("input.json")
	if err != nil {
		panic(err.Error())
	}
	jsonToByteArray, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}
	var cs Customers
	json.Unmarshal(jsonToByteArray, &cs)

	for i := 0; i < len(cs.Customers); i++ {
		var newCustomer models.Customer

		newCustomer.Name = cs.Customers[i].Name
		newCustomer.Address = cs.Customers[i].Address
		newCustomer.Age = cs.Customers[i].Age
		newCustomer.Email = cs.Customers[i].Email
		// fmt.Println(newCustomer)
		// db.InsertCustomer(newCustomer)
	}
	var c models.Customer
	c.Name = "foo"
	c.Address = "999 west"
	c.Age = 99
	c.Email = "bar@gmail.com"

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
	
	`)

}
func add(w http.ResponseWriter, r *http.Request) {
	var c models.Customer
	d := db.InitDB()
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
	d := db.InitDB()
	var arr []models.Customer = db.GetAllCustomers(d)
	for _, p := range arr {
		fmt.Fprint(w, p)
	}
}
func GetOneCustomer(w http.ResponseWriter, r *http.Request) {
	d := db.InitDB()
	customerEmail := mux.Vars(r)["email"]
	c, err := db.GetSpecificCustomer(d, customerEmail)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, c)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	db.InitDB()
	customerEmail := mux.Vars(r)["email"]
	db.RemoveSpecificCustomer(customerEmail)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {

}
