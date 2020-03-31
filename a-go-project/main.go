package main

import (
	"a-go-project/db"
	"a-go-project/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Customers struct {
	Customers []models.Customer `json:"Customers"`
}

func main() {
	db.InitDB()

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

		//INSERT here or somewhere else?
	}

}
