package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	//After adding testutils_test.go file
	app := newTestApplication(t)
	tlsServer := newTestServer(t, app.routes())
	defer tlsServer.Close()

	//Lets get the status code, header, and body
	statusCode, header, body := tlsServer.get(t, "/ping")

	if statusCode != http.StatusOK {
		t.Errorf("\nGot: %q\n Want: %d", statusCode, http.StatusOK)
	}
	fmt.Println(header)

	if string(body) != "OK" {
		t.Errorf("\nGot: %q\n Want: OK", body)
	}

}

//End to end
func TestE2EPing(t *testing.T) {
	//Create new app struct and mock the loggers that will discard anything written to them
	app := &application{
		InfoLog:  log.New(ioutil.Discard, "", 0),
		errorLog: log.New(ioutil.Discard, "", 0),
	}

	// We then use the httptest.NewTLSServer() function to create a new test
	// server, passing in the value returned by our app.routes() method as the
	// handler for the server. This starts up a HTTPS server which listens on a
	// randomly-chosen port of your local machine for the duration of the test.
	// Notice that we defer a call to ts.Close() to shutdown the server when
	// the test finishes.
	tlsServer := httptest.NewTLSServer(app.routes())
	defer tlsServer.Close()

	// The network address that the test server is listening on is contained
	// in the ts.URL field. We can use this along with the ts.Client().Get()
	// method to make a GET /ping request against the test server. This
	// returns a http.Response struct containing the response.
	result, err := tlsServer.Client().Get(tlsServer.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	//check the status code
	if result.StatusCode != http.StatusOK {
		t.Errorf("\nGot: %q \nWant: %q", result.StatusCode, http.StatusOK)
	}

	//check body
	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("\nGot: %q\nWant: OK", body)
	}
}

func TestShowSnippet(t *testing.T) {
	// Create a new instance of our application struct which uses the mocked // dependencies.
	app := newTestApplication(t)
	// Establish a new test server for running end-to-end tests.
	tlsServer := newTestServer(t, app.routes())
	defer tlsServer.Close()
	// Set up some table-driven tests to check the responses sent by our // application for different URLs.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("Mock")},
		{"Non-existent ID", "/snippet/2", http.StatusNotFound, nil},
		{"Negative ID", "/snippet/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/snippet/1.23", http.StatusNotFound, nil},
		{"String ID", "/snippet/foo", http.StatusNotFound, nil},
		{"Empty ID", "/snippet/", http.StatusNotFound, nil},
		{"Trailing slash", "/snippet/1/", http.StatusNotFound, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			code, _, body := tlsServer.get(t, test.urlPath)
			fmt.Println(code)
			fmt.Println(test.wantCode)
			fmt.Println(body)
			fmt.Println(test.wantBody)
			if code != test.wantCode {

				t.Errorf("\nGot: %q\nWant: %q", code, test.wantCode)
			}

			if !bytes.Contains(body, test.wantBody) {
				t.Errorf("want body to contain %q", test.wantBody)
			}
		})
	}
}

func TestSignUp(t *testing.T) {
	// Create a new instance of our application struct which uses the mocked // dependencies.
	app := newTestApplication(t)
	// Establish a new test server for running end-to-end tests.
	tlsServer := newTestServer(t, app.routes())
	defer tlsServer.Close()

	//make the get request to the signup user
	_, _, body := tlsServer.get(t, "/user/signup")

	token := getCSRFToken(t, body)

	t.Log(token)
}
