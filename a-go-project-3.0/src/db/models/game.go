package models

import (
	"errors"
	"time"
)

//ErrNoGameFound - error if no game is found when checking DB
var ErrNoGameFound = errors.New("no game found")

//Game -
type Game struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"releaseDate"`
	Platform    string    `json:"platform"`
	Genre       string    `json:"genre"`
	Rating      int       `json:"rating"`
}
