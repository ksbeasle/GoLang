package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ksbeasle/GoLang/api"
	"github.com/ksbeasle/GoLang/application"
)

func main() {

	/* Start a Connection to the Database - mysql */
	DB, err := application.StartDB()
	app := &application.app{
		GameDB: DB,
	}
	defer DB.Close()

	// if err != nil {
	// 	errorLog.Fatal(err)
	// }
	//defer DB.Close()
	// /* Create a new application struct */
	// app := &application{
	// 	infoLog:  infoLog,
	// 	errorLog: errorLog,
	// 	db:       &mysql.DBModel{DB: DB},
	// 	//gonna check something here later -- db:       DB,
	// }

	/* server struct */
	server := &http.Server{
		Addr:    ":8080",
		Handler: api.Routes(),
	}

	infoLog.Println("STARTING SERVER AT PORT ... ", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
