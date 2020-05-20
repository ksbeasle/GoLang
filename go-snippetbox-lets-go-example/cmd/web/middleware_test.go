package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	//response recorder
	responseRecorder := httptest.NewRecorder()

	//dummy request
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP handler that we can pass to our secureHeaders
	// middleware, which writes a 200 status code and "OK" response body.
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Pass the mock HTTP handler to our secureHeaders middleware. Because
	// secureHeaders *returns* a http.Handler we can call its ServeHTTP()
	// method, passing in the http.ResponseRecorder and dummy http.Request to
	// execute it.
	secureHeaders(mockHandler).ServeHTTP(responseRecorder, req)

	//get the result of the call to the method
	result := responseRecorder.Result()

	//check that frames headers were set properly
	frame := result.Header.Get("X-Frame-Options")
	if frame != "deny" {
		t.Errorf("\nGot: %q\nWant: deny", frame)
	}

	//check X-XSS-Protection
	xss := result.Header.Get("X-XSS-Protection")
	if xss != "1; mode=block" {
		t.Errorf("\nGot: %q\nWant: 1; mode=block", xss)
	}

	//check that the middleware called the next handler correctly
	if result.StatusCode != http.StatusOK {
		t.Errorf("\nGot: %q\nWant: %q", result.StatusCode, http.StatusOK)
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
