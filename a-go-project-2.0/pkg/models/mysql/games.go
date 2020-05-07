package mysql

import (
	"database/sql"

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

	}
	return games, nil
}
