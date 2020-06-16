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
