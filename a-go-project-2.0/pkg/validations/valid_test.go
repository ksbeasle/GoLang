package validations

import (
	"testing"
)

func TestValidTitle(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		wantError error
	}{
		{"Valid Title", "title", nil},
		{"Invalid Title", "", ErrEmptyTitle},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidTitle(tc.title)

			if err != tc.wantError {
				t.Errorf("Got: %v\nWant: %v", err, tc.wantError)
			}

		})
	}
}

func TestValidGenre(t *testing.T) {

	tests := []struct {
		name      string
		genre     string
		wantError error
	}{
		{"Valid Genre", "genre", nil},
		{"Invalid Genre", "", ErrEmptyGenre},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidGenre(tc.genre)

			if err != tc.wantError {
				t.Errorf("Got: %v\nWant: %v", err, tc.wantError)
			}
		})
	}
}

func TestValidRating(t *testing.T) {
	tests := []struct {
		name      string
		rating    int
		wantError error
	}{
		{"Valid Rating", 10, nil},
		{"Invalid Rating", 0, ErrInvalidRating},
		{"Invalid Rating negative", -1, ErrInvalidRating},
		{"Invalid Rating out of bounds", 100, ErrInvalidRating},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidRating(tc.rating)

			if err != tc.wantError {
				t.Errorf("Got: %v\nWant: %v", err, tc.wantError)
			}
		})
	}
}
