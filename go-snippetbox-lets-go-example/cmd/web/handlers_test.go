package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	//After adding testutils_test.go file
	app := newTestApplicaion()
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
