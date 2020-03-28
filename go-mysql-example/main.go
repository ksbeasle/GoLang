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
	//TODO: apparently theres no password for 'root' I should fix that
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//INSERT into db
	insert, err := db.Query("INSERT INTO test.user VALUES (4, 'grom4m')")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	//SELECT
	res, err := db.Query("SELECT ID, Name from user")
	if err != nil {
		panic(err.Error())
	}
	for res.Next() {
		var e example
		err = res.Scan(&e.ID, &e.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(e.Name)
	}

	//SELECT SINGLE ROW
	var e example
	err = db.QueryRow("SELECT ID, Name from user WHERE ID = ?", 2).Scan(&e.ID, &e.Name)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("------------")
	fmt.Println(e.ID)
	fmt.Println(e.Name)
}
