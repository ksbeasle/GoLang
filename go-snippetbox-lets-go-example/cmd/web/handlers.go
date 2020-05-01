package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//http://localhost:8080/snippet?id=123

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	//this will prevent the slash catching all requests
	if r.URL.Path != "/" {
		//http.NotFound(w, r)
		app.NotFound(w)
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//using application error log instead of the default log
		//app.errorLog.Println(err.Error())
		app.ServerError(w, err)
		//http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		//app.errorLog.Println(err.Error())
		//http.Error(w, "Internal server error", http.StatusInternalServerError)
		app.ServerError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		//app.errorLog.Println(err.Error())
		//http.Error(w, "Invalid ID", http.StatusNotFound)
		//http.NotFound(w,r)
		app.NotFound(w)
		return
	}
	fmt.Fprintf(w, "Id = %d", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	//Allow only POST method to be used per HTTP good practices
	if r.Method != http.MethodPost {
		//Let user know what Methods are accepted at the endpoint
		w.Header().Set("Allow", http.MethodPost)
		app.ClientError(w, http.StatusMethodNotAllowed)
		//app.errorLog.Println(http.StatusMethodNotAllowed)
		//405
		//http.Error(w, "Method Now allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a snippet"))
}
