package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ksbeasle/GoLang/pkg/models"
	"github.com/ksbeasle/GoLang/pkg/models/mysql"
)

/*Run the command in this comment to run app
***************IMPORTANT******************
******************************************
go run $(ls -1 *.go | grep -v _test.go)
******************************************
******************************************
*/
//Dependencies for use across the entire application
//We changed vgmodel to an interface in order to use mocks for testing
//As long as those methods are satisfied everything should run fine
type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	vgmodel  interface {
		Insert(title string, genre string, rating int, platform string, releaseDate string) (int, error)
		All() ([]*models.Game, error)
		Get(id int) (*models.Game, error)
	}
}

func main() {
	//logs
	infoLog := log.New(os.Stdout, "info: ", log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stdout, "error: ", log.Ltime|log.Lshortfile)

	//Initialize Database
	db, err := startDB()
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	//Dependencies
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		vgmodel:  &mysql.VGModel{DB: db},
	}
	//Created this struct for a cleaner look
	serve := &http.Server{
		ErrorLog: errorLog,
		Addr:     ":8080",
		Handler:  app.routes(),
	}
	//Start server
	infoLog.Println("STARTING SERVER ON PORT 8080: ...")
	err = serve.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}

}

//Database - Connect to DB
func startDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "web1:pass@tcp(localhost:3306)/videogames")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("DATABASE SUCCESSFULLY CONNECTED")
	return db, nil
}
