package database

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ksbeasle/GoLang/database/models"
)

type GameDB struct {
	GameDB *sql.DB
}

/*startDB - Connect to the mysql database, return the db if successful else an error */
func StartDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "web3:pass@tcp(localhost:3306)/videogames3?parseTime=true")
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

//func All()

/*Get - Will get a specific game based on the id else return an error */
func (g *GameDB) Get(id int) (*models.Game, error) {
	//connect to db
	//db, err := StartDB()
	if err != nil {
		log.Println("Unable to connect to DB: ", err)
		return nil, err
	}

	//close db
	//defer db.Close()

	//query statement
	stmt := `SELECT title, description, releaseDate, platform, genre, rating
			 FROM games
			 WHERE id=?`

	//execute statement
	row := g.QueryRow(stmt, id)

	//game to hold the values from the query
	game := &models.Game{}

	//scan the values into g
	err = row.Scan(&game.Title, &game.Description, &game.ReleaseDate, &game.Platform, &game.Genre, &game.Rating)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoGameFound
		} else {
			return nil, err
		}
	}

	return game, nil
}

/*Insert - */
func Insert(title string, desc string, release string, platform string, genre string, rating int) (int, string, error) {
	return 0, "", nil
}
