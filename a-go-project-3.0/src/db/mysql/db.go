package mysql

import (
	"database/sql"
)

type DBModel struct {
	DB *sql.DB
}

func (d *DBModel) Insert(title string, desc string, release string, platform string, genre string, rating int) (int, error, string) {
	return nil, nil, nil
}
