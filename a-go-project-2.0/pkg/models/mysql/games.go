package mysql

import "database/sql"

//This struct will act on the database
type VGModel struct {
	DB *sql.DB
}

//Insert a new video game
func (v *VGModel) Insert(title string, genre string, rating int, platform string, releaseDate string) (int, error) {
	//Query
	stmt := `INSERT INTO games (title, genre, rating, platform, releaseDate)
			VALUES (?, ?, ?, ?, ?)`
	//Prepare DB statement
	res, err := v.DB.Exec(stmt, title, genre, rating, platform, releaseDate)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	//Convert id to int, since LastInsertId() returns int64
	return int(id), nil
}
