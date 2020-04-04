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
		//fmt.Println(newCustomer)
		//db.InsertCustomer(newCustomer)
	}
	var c models.Customer
	c.Name = "foo"
	c.Address = "999 west"
	c.Age = 99
	c.Email = "bar@gmail.com"

	//db.InsertCustomer(c)
	db.GetAllCustomers()

}
