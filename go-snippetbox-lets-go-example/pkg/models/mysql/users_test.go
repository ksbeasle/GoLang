package mysql

import (
	"reflect"
	"testing"
	"time"

	"ksbeasle.net/snippetbox/pkg/models"
)

func TestUserModelGet(t *testing.T) {
	//Skip test if Short()
	if testing.Short() {
		t.Skip("User Model Get Test -- skipping")
	}

	tests := []struct {
		name      string
		userID    int
		wantUser  *models.User
		wantError error
	}{
		{
			name:   "Valid ID",
			userID: 1,
			wantUser: &models.User{
				ID:      1,
				Name:    "user name",
				Email:   "email@example.com",
				Created: time.Date(2018, 12, 23, 17, 25, 22, 0, time.UTC),
				Active:  true,
			},
		},
		{
			name:      "Zero ID",
			userID:    0,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
		{
			name:      "ID doesn't Exist",
			userID:    2,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			//Create new DB connection and teardown func() to defer after testing is complete
			testdb, teardown := newTestDB(t)

			defer teardown()

			//create instance of usermodel
			um := UserModel{testdb}

			//Get the user
			user, err := um.Get(tc.userID)
			if err != tc.wantError {
				t.Errorf("\nWant: %v\nGot: %s", tc.wantError, err)
			}

			//Using the reflect.DeepEqual() function is an effective way to check for
			// equality between arbitrarily complex custom types.

			if !reflect.DeepEqual(tc.wantUser, user) {
				t.Errorf("\nWant: %v\nGot: %v", tc.wantUser, user)
			}

		})
	}
}
