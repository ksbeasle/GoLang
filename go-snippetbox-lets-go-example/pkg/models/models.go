package models

import (
	"errors"
	"time"
)

var errNoRecord = errors.New("No record found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	created time.Time
	Expires time.Time
}
