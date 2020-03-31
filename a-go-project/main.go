package main

import (
	"a-go-project/db"
	"a-go-project/models"
	"encoding/json"
	"fmt"
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
		fmt.Println(cs.Customers[i])
	}

}
