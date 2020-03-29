package main

import (
	"a-go-project/models"
	"fmt"
)

func main() {
	fmt.Print("yo")
	m := make(map[string]interface{})
	m["ok"] = "ok"
	fmt.Println(m)
	c := models.Customer{}
	c.Age = 2

}
