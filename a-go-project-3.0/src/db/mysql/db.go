package mysql

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ksbeasle/GoLang/db/models"
)

/*startDB - Connect to the mysql database, return the db if successful else an error */
func startDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "web3:pass@tcp(localhost:3306)/videogames3")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("CONNCETION TO DATABASE SUCCESSFUL")
	return db, nil
}

/*Get - Will get a specific game based on the id else return an error */
func Get(id int) (*models.Game, error) {
	//connect to db
	db, err := startDB()
	if err != nil {
		log.Println("Unable to connect to DB: ", err)
		return nil, err
	}

	//close db
	defer db.Close()

	//query statement
	/*
		`SELECT title, description, releaseDate, platform, genre, rating
				FROM games
			 	WHERE id=?`
	*/
	stmt := `SELECT title, description, platform, genre, rating
			FROM games
		 	WHERE id=?`

	//execute statement
	row := db.QueryRow(stmt, id)

	//game to hold the values from the query
	g := &models.Game{}

	//scan the values into g
	err = row.Scan(&g.Title, &g.Description, &g.Platform, &g.Genre, &g.Rating)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoGameFound
		} else {
			return nil, err
		}
	}

	return g, nil
}

/*Insert - */
func Insert(title string, desc string, release string, platform string, genre string, rating int) (int, string, error) {
	return 0, "", nil
}
