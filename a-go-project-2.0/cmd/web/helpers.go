package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

//Client side Error
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

//Server side Error
func (app *application) serverError(w http.ResponseWriter, err error) {
	log.Println("IT here")
	//Using stacktrace so we can get more information on the server side error for debugging
	stacktrace := fmt.Sprintf("%s \n %s", err.Error(), debug.Stack())

	//We want to see the actual line where the error occurred so we change the frame depth to 2
	app.errorLog.Output(2, stacktrace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

/* Convenient wrapper for ClientError to send 404 not found response */
func (app *application) NotFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
