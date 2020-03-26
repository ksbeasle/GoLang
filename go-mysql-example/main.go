package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type example struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//INSERT into db
	insert, err := db.Query("INSERT INTO user VALUES (2 `gromm`)")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	//SELECT
	res, err := db.Query("SELECT ID, Name from user")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res)
	fmt.Println(err)
	for res.Next() {
		var e example
		err = res.Scan(&e.ID, &e.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(e.Name)
	}
}
