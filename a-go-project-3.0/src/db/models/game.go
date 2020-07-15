package models

import (
	"errors"
)

//Error - No game found
var errNoGameFound = errors.New("no game found")

//add json l8r `json:blah blah`
type Game struct {
	uuid        int
	Title       string
	Description string
	ReleaseDate string
	Platform    string
	Genre       string
	Rating      int
}
