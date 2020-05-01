package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

/* Server error log */
func (app *application) ServerError(w http.ResponseWriter, err error) {
	stackTrace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	//app.errorLog.Println(stackTrace)
	//We want to see the actual line where this error has occurred by changing errorLog's frame depth to 2
	app.errorLog.Output(2, stackTrace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

/* Client error log */
func (app *application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

/* Convenient wrapper for ClientError to send 404 not found response */
func (app *application) NotFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
