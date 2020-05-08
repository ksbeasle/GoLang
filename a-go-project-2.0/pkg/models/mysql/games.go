package mysql

import (
	"database/sql"
	"errors"

	"github.com/ksbeasle/GoLang/pkg/models"
)

//This struct will act on the database
type VGModel struct {
	DB *sql.DB
}

//Insert a new video game
func (v *VGModel) Insert(title string, genre string, rating int, platform string, releaseDate string) (int, error) {
	//Query
	stmt := `INSERT INTO games (title, genre, rating, platform, releaseDate)
			VALUES (?, ?, ?, ?, ?)`
	//Execute DB statement
	res, err := v.DB.Exec(stmt, title, genre, rating, platform, releaseDate)
	if err != nil {
		return 0, err
	}
	//Get the id of the inserted game
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	//Convert id to int, since LastInsertId() returns int64
	return int(id), nil
}

//Get all games
func (v *VGModel) All() ([]*models.Game, error) {
	//Query
	stmt := `SELECT id, title, genre, rating, platform, releaseDate
			 FROM games`
	//Execute DB statement
	rows, err := v.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	//Ensure we close resultset or a possible Query() error can occur on a nil resultset
	defer rows.Close()

	//Slice to hold the list of games
	games := []*models.Game{}

	//loop through each row
	for rows.Next() {
		//new struct to hold the values of the row
		g := &models.Game{}
		//Scan the row into the newly made struct
		err := rows.Scan(&g.ID, &g.Title, &g.Genre, &g.Rating, &g.Platform, &g.ReleaseDate)
		if err != nil {
			return nil, err
		}

		//append the games array
		games = append(games, g)

	}

	//Check one more time for any rows errors since an error
	//can still occur even after the for loop
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

//GET - get a single game from given ID
func (v *VGModel) Get(id int) (*models.Game, error) {
	//Query
	stmt := `SELECT id, title, genre, rating, platform, releaseDate
			 FROM games
			 WHERE id = ?`
	//execute query
	row := v.DB.QueryRow(stmt, id)

	//new struct to hold row values
	g := &models.Game{}

	//Scan values into the new Game struct
	err := row.Scan(&g.ID, &g.Title, &g.Genre, &g.Rating, &g.Platform, &g.ReleaseDate)

	//This error check is a little different here we are checking if there was no game
	//found from the given ID
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoGameFound
		} else {
			return nil, err
		}
	}

	return g, nil
}
