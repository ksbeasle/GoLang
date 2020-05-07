package models

import "errors"

var errNoGameFound = errors.New("No game found.")

type Game struct {
	ID          int
	Title       string
	Genre       string
	Rating      int
	Platform    string
	ReleaseDate string
}
