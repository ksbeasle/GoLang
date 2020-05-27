package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ksbeasle/GoLang/pkg/models"
	v "github.com/ksbeasle/GoLang/pkg/validations"
)

//HOME
func (app *application) index(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("HOME ... ")
	if r.URL.Path != "/" {
		app.NotFound(w)
	}
	g, err := app.vgmodel.All()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, game := range g {
		fmt.Fprintf(w, "%v\n", game)
	}

}

//Add a game
func (app *application) add(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("ADDING GAME ... ")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	//Game to save the data in
	var g = models.Game{}
	err = json.Unmarshal(reqBody, &g)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	//Validate user input before inserting
	err = v.ValidTitle(g.Title)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	err = v.ValidGenre(g.Genre)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	err = v.ValidRating(g.Rating)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	err = v.ValidPlatform(g.Platform)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	err = v.ValidReleaseDate(g.ReleaseDate)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	id, err := app.vgmodel.Insert(g.Title, g.Genre, g.Rating, g.Platform, g.ReleaseDate)
	if err != nil {
		app.serverError(w, err)
	}
	fmt.Fprintf(w, "%d successfully created", id)
}

//Get a single game based on given id
func (app *application) getGame(w http.ResponseWriter, r *http.Request) {
	//Get the id passed in from user (URL)
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	g, err := app.vgmodel.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoGameFound) {
			app.NotFound(w)
		} else {
			app.serverError(w, err)
		}
	}

	fmt.Fprintf(w, "%v", g)

}
