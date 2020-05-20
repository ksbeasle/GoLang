package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{
			name: "UTC",
			t:    time.Date(2000, 1, 01, 0, 0, 0, 0, time.UTC),
			want: "01 Jan 2000 at 00:00",
		},
		{
			name: "Empty",
			t:    time.Time{},
			want: "",
		},
		{
			name: "CET",
			t:    time.Date(2020, 12, 17, 10, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Dec 2020 at 10:00",
		},
	}

	for _, testCase := range tests {
		// Use the t.Run() function to run a sub-test for each test case.
		//The first parameter to this is the name of the test (which is used to
		// identify the sub-test in any log output) and the second parameter is
		// and anonymous function containing the actual test for each case.
		t.Run(testCase.name, func(t *testing.T) {
			date := humanDate(testCase.t)

			if date != testCase.want {
				t.Errorf("\nWant: %q\nGot: %q", testCase.want, date)
			}
		})
	}

}
