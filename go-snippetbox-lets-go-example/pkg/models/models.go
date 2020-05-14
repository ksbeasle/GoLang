package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("No record found")

	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrDuplicateEmail = errors.New("Duplicate email")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
