package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ksbeasle/GoLang/pkg/models/mock"
)

//NewTestApplication that contains the dependencies: logs will be discarded and mocks will be used for games
func NewTestApplication(t *testing.T) *application {
	return &application{
		infoLog:  log.New(ioutil.Discard, "", 0),
		errorLog: log.New(ioutil.Discard, "", 0),
		games:    &mock.VGModel,
	}
}

type testServer struct {
	*httptest.Server
}

//Test server
func NewTestServer(t *testing.T, hanlder http.Handler) *testServer {
	server := httptest.NewServer(handler)
	return &testServer{server}
}

//This get() will mock a GET request against the test server by passing in a url
// and returns the status code and body
func (ts *testServer) get(t *testing.T, url string) (int, []byte) {
	result, err := ts.Client().Get(ts.URL + url)
	if err != nil {
		t.Fatal(err)
	}

	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	return result.StatusCode, body
}
