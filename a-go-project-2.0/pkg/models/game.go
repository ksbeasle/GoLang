package models

import "errors"

//Error if no game was found
var ErrNoGameFound = errors.New("No game found.")

//Game struct
type Game struct {
	ID          int
	Title       string
	Genre       string
	Rating      int
	Platform    string
	ReleaseDate string
}
