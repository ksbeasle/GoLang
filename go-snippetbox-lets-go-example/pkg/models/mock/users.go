package mock

import (
	"time"

	"ksbeasle.net/snippetbox/pkg/models"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "user name",
	Email:   "email@example.com",
	Created: time.Now(),
	Active:  true,
}

type UserModel struct{}

func (um *UserModel) Insert(name string, email string, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (um *UserModel) Authenticate(email string, password string) (int, error) {
	switch email {
	case "email@example.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (um *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
