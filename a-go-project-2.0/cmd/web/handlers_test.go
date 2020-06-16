package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetGame(t *testing.T) {

	//Test application
	app := NewTestApplication(t)

	//Test server
	server := NewTestServer(t, app.routes())
	defer server.Close()

	testCases := []struct {
		name     string
		url      string
		wantCode int
		wantBody []byte
	}{
		{"Valid Request", "/game/1", http.StatusOK, []byte("Halo 3")},
		{"Negative ID", "/game/-1", http.StatusNotFound, nil},
		{"ID doesn't Exists", "/game/0", http.StatusNotFound, nil},
		{"ID is a String", "/game/str", http.StatusNotFound, nil},
		{"ID is a decimal", "/game/1.0", http.StatusNotFound, nil},
		{"ID is empty", "/game/", http.StatusNotFound, nil},
		{"Trailing slash", "/game/1/", http.StatusNotFound, nil},
	}

	//loop through testCases
	for _, tc := range testCases {
		code, body := server.get(t, tc.url)
		//check if the status code matches
		if code != tc.wantCode {

			t.Errorf("\nGot: %q\nWant: %d", code, tc.wantCode)
		}

		//HOW DOES !bytes.Contains(tc.wantBody, body) affect the code test case????
		if !bytes.Contains(body, tc.wantBody) {
			t.Errorf("\nGot: %q\nWant: %q", body, tc.wantBody)
		}
	}

}
