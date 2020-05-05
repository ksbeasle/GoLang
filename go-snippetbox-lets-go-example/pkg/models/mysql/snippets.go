package mysql

import (
	"database/sql"
	"errors"
	"log"

	"ksbeasle.net/snippetbox/pkg/models"
)

//snippetodel type that wraps a sql.Db
//we will inject this in main.go
type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	log.Println("INSERT ...")
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

	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Snippet{}

	//row.scan to copy each value form sql.row
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires 
			 FROM snippets 
			 WHERE expires > UTC_TIMESTAMP()
			 ORDER BY created
			 DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	//we defer after error check to ensure sql resultset is closed
	//otherwise you can get a Query() error and a panic will happen when trying to close nil resultset
	defer rows.Close()

	//slice to hold results
	snippets := []*models.Snippet{}

	//iterate over rows
	for rows.Next() {
		//pointer to new zeroed snippet struct
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	//Important to call rows.Err() because an error can still occur even
	// after the for loop finishes
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
