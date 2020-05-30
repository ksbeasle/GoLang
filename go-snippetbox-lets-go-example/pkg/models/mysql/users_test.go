package mysql

import (
	"go-snippetbox-lets-go-example/pkg/models"
	"testing"
	"time"
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
		{name: "Valid ID",
			userID: 1,
			wantUser: &models.User{
				ID:      1,
				Name:    "user name",
				Email:   "email@example.com",
				Created: time.Now(),
				Active:  true,
			}},
	}
}
