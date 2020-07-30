package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ksbeasle/GoLang/api"
	"github.com/ksbeasle/GoLang/application"
	"github.com/ksbeasle/GoLang/database"
)

func main() {

	/* Start a Connection to the Database - mysql */
	db, err := application.StartDB()
	if err != nil {
		log.Fatal(err)
	}

	app := &application.App{
		DBMODEL: &database.GameDB{DB: db},
	}
	log.Println(app)
	defer db.Close()

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

	log.Println("STARTING SERVER AT PORT ... ", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
