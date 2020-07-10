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
		{"Invalid Title", "", errEmptyTitle},
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
		{"Invalid Genre", "", errEmptyGenre},
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
		{"Invalid Rating", 0, errInvalidRating},
		{"Invalid Rating negative", -1, errInvalidRating},
		{"Invalid Rating out of bounds", 100, errInvalidRating},
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

func TestValidateReleaseDate(t *testing.T) {
	tests := []struct {
		name      string
		date      string
		wantError error
	}{
		{"Valid Date", "October 26, 1995", nil},
		{"Valid Date 2", "November 26, 1995", nil},
		{"Valid Date Feb", "February 28, 1995", nil},
		{"Invalid Date - February 30th", "February 30, 1995", errInvalidDayFeb},
		{"Invalid Date - September 31st", "September 31, 2000", errInvalidDay30},
		{"Invalid Date empty", "", errEmptyReleaseDate},
		{"Invalid Date bad day", "December 32, 1995", errInvalidDay},
		{"Invalid Date bad day for particular month", "February 31, 1995", errInvalidDayFeb},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidReleaseDate(tc.date)

			if err != tc.wantError {
				t.Errorf("Got: %v\nWant: %v", err, tc.wantError)
			}
		})
	}
}
