package main

import (
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/golangcollege/sessions"
	"ksbeasle.net/snippetbox/pkg/models/mock"
)

//Application struct that returns mocked dependencies
func newTestApplication(t *testing.T) *application {

	tc, err := newTemplateCache("./../../ui/html/")
	if err != nil {
		t.Fatal(err)
	}

	// Create a session manager instance, with the same settings as production.
	session := sessions.New([]byte("3dSm5MnygFHh7XidAtbskXrjbwfoJcbJ"))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	return &application{
		InfoLog:       log.New(ioutil.Discard, "", 0),
		errorLog:      log.New(ioutil.Discard, "", 0),
		session:       session,
		templateCache: tc,
		snippets:      &mock.SnippetModel{},
		users:         &mock.UserModel{},
	}
}

//test server that anonymously embeds httptest.server
type testServer struct {
	*httptest.Server
}

//Returns a new test server
func newTestServer(t *testing.T, handler http.Handler) *testServer {
	tlsServer := httptest.NewTLSServer(handler)

	//create a cookie so that we can test anti-csrf cases
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	//add cookie to client
	tlsServer.Client().Jar = jar

	//disable any redirect requests
	//instead return the receieved response
	tlsServer.Client().CheckRedirect = func(r *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{tlsServer}
}

// Implement a get method on our custom testServer type. This makes a GET
// request to a given url path on the test server, and returns the response
// status code, headers and body.
func (tlsServer *testServer) get(t *testing.T, url string) (int, http.Header, []byte) {
	res, err := tlsServer.Client().Get(tlsServer.URL + url)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	return res.StatusCode, res.Header, body
}

//regular expression which captures CSRF token
var csrfToken = regexp.MustCompile(`<input type='hidden' name='csrf_token' value='(.+)'>`)

func getCSRFToken(t *testing.T, body []byte) string {
	//using FindSubmatch method to take the token from the html body
	match := csrfToken.FindSubmatch(body)
	if len(match) < 2 {
		t.Fatal("no CSRF token found")
	}
	return html.UnescapeString(string(match[1]))
}

// func (ts *testServer) postForm(t *testing.T, url string, form url.Values) (int, http.Header, []byte) {
// 	rs, err := ts.Client().PostForm(ts.URL+url, form) if err != nil {
// 		t.Fatal(err) }
// 		// Read the response body.
// 		defer rs.Body.Close()
// 		body, err := ioutil.ReadAll(rs.Body) if err != nil {
// 		t.Fatal(err) }
// 		// Return the response status, headers and body.
// 		return rs.StatusCode, rs.Header, body
// }
