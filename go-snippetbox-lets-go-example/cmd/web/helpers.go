package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
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

//Render templates from the Cache to avoid duplicated code
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.ServerError(w, fmt.Errorf("%s does not exist", name))
		return
	}
	//new buffer
	buf := new(bytes.Buffer)

	//we are going to write to our buffer first to check for any errors
	//before writing it to the client
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.ServerError(w, err)
		return
	}
	td.Flash = app.session.PopString(r, "flash")
	// Write the contents of the buffer to the http.ResponseWriter. Again, this
	// is another time where we pass our http.ResponseWriter to a function that
	// takes an io.Writer.
	buf.WriteTo(w)
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}
