package mysql

import (
	"database/sql"
	"log"

	"ksbeasle.net/snippetbox/pkg/models"
)

//snippetodel type that wraps a sql.Db
//we will inject this in main.go
type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	log.Println()
	stmt := `INSERT INTO snippets (title, content, created, expires) 
	VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(),INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, content, expires, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	//LastInsertId returns int64 so we convert it to int
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
