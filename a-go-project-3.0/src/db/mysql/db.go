package mysql

import (
	"database/sql"

	"github.com/ksbeasle/GoLang/db/models"
)

type DBModel struct {
	DB *sql.DB
}

/* Get - */
func Get(id int) (*models.Game, error) {
	g := &models.Game{}
	return g, nil
}

/* Insert - */
func (d *DBModel) Insert(title string, desc string, release string, platform string, genre string, rating int) (int, error, string) {
	return 0, nil, ""
}
