package models

import (
	"errors"
	"time"
)

//Error - No game found
var errNoGameFound = errors.New("no game found")

type Game struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"releaseDate"`
	Platform    string    `json:"platform"`
	Genre       string    `json:"genre"`
	Rating      int       `json:"rating"`
}
