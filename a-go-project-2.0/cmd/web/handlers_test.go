package main

import (
	"net/http"
	"testing"
)

func TestGetGame(t *testing.T) {

	//Test server
	server := NewTestServer(t, app.routes())
	defer server.Close()
	//Test application
	app := NewTestApplication(t)

	testCases := []struct {
		name     string
		url      string
		wantBody []byte
		wantCode int
	}{
		{"Valid Request", "/game/1", []byte("Halo 3"), http.StatusOK},
		{"Negative ID", "/game/-1", nil, http.StatusNotFound},
		{"ID doesn't Exists", "/game/0", nil, http.StatusNotFound},
		{"ID is a String", "/game/str", nil, http.StatusNotFound},
		{"ID is a decimal", "/game/1.0", nil, http.StatusNotFound},
		{"ID is empty", "/game/", nil, http.StatusNotFound},
		{"Trailing slash", "/game/1/", nil, http.StatusNotFound},
	}

	//loop through testCases
	for _, tc := range testCases {
		code, body := server.get(t, tc.url)

		//check if the status code matches
		if code != tc.wantCode {
			t.Errorf("\nGot: %d\nWant: %d", code, http.StatusNotFound)
		}

		if 
	}

}