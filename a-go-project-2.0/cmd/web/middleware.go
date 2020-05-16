package main

import "net/http"

//This middleware will log all incoming requests
func (app *application) LogIncomingRequests(handler http.Handler) http.Handler {
	//closure so that we can have access to the handler that is passed in
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//make our log from the incoming request
		app.infoLog.Printf("\nMETHOD: %s \n - URL: %v \n - HOST: %s \n - HEADER: %v", r.Method, r.URL.RequestURI(), r.Host, r.Header)
		//pass the handler
		handler.ServeHTTP(w, r)
		//if we wanted to do anything after the chain of hanlders we would do so after the ServeHttp()
	})
}
