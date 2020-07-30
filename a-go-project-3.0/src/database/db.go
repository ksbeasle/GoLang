package database

import (
	"database/sql"
	"errors"

	"github.com/ksbeasle/GoLang/database/models"
)

type GameDB struct {
	DB *sql.DB
}

//func All()

/*Get - Will get a specific game based on the id else return an error */
func (g *GameDB) Get(id int) (*models.Game, error) {
	//query statement
	stmt := `SELECT title, description, releaseDate, platform, genre, rating
			 FROM games
			 WHERE id=?`

	//execute statement
	row := g.DB.QueryRow(stmt, id)

	//game to hold the values from the query
	game := &models.Game{}

	//scan the values into g
	err := row.Scan(&game.Title, &game.Description, &game.ReleaseDate, &game.Platform, &game.Genre, &game.Rating)

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
