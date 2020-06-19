package mysql

import (
	"reflect"
	"testing"

	"github.com/ksbeasle/GoLang/pkg/models"
)

func TestGet(t *testing.T) {

	tests := []struct {
		Name      string
		GameID    int
		wantGame  *models.Game
		wantError error
	}{
		{
			Name:   "Valid ID",
			GameID: 1,
			wantGame: &models.Game{
				ID:          1,
				Title:       "Game Title",
				Genre:       "Game Genre",
				Rating:      1,
				Platform:    "Game Platform",
				ReleaseDate: "January 1, 2000",
			},
		},
		{
			Name:      "Invalid ID",
			GameID:    0,
			wantGame:  nil,
			wantError: models.ErrNoGameFound,
		},
		{
			Name:      "ID doesn't Exist",
			GameID:    2,
			wantGame:  nil,
			wantError: models.ErrNoGameFound,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.Name, func(t *testing.T) {
			db, teardownDB := createTestDB(t)

			//Close DB once we are done
			defer teardownDB()

			v := VGModel{db}

			//Get the game
			game, err := v.Get(testcase.GameID)
			if err != testcase.wantError {
				t.Errorf("\nGot: %s\nWant: %v", err, testcase.wantError)
			}

			if !reflect.DeepEqual(testcase.wantGame, game) {
				t.Errorf("\nGot: %v\nWant: %v", game, testcase.wantGame)
			}
		})
	}
}

func TestInsert(t *testing.T) {

	tests := []struct {
		Name         string
		GameID       int
		InsertedGame *models.Game
		wantError    error
		wantId       int
	}{
		{
			Name:   "Valid Insert",
			GameID: 1,
			InsertedGame: &models.Game{
				ID:          1,
				Title:       "Test",
				Genre:       "Test",
				Rating:      10,
				Platform:    "Test",
				ReleaseDate: "January 1, 2000",
			},
			wantError: nil,
			wantId:    2,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.Name, func(t *testing.T) {
			db, teardownDB := createTestDB(t)

			defer teardownDB()

			v := VGModel{db}
			//title string, genre string, rating int, platform string, releaseDate string
			id, err := v.Insert(testcase.InsertedGame.Title, testcase.InsertedGame.Genre, testcase.InsertedGame.Rating, testcase.InsertedGame.Platform, testcase.InsertedGame.ReleaseDate)

			if id != testcase.wantId {
				t.Errorf("\nGot: %d\nWant: %v", id, testcase.wantId)
			}

			if err != testcase.wantError {
				t.Errorf("\nGot: %s\nWant: %v", err, testcase.wantError)
			}
		})
	}

}
