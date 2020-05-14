package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"ksbeasle.net/snippetbox/pkg/forms"
	"ksbeasle.net/snippetbox/pkg/models"
)

//http://localhost:8080/snippet?id=123

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	//this will prevent the slash catching all requests
	//third party package pat catches the "/" so no need for this check
	// if r.URL.Path != "/" {
	// 	//http.NotFound(w, r)
	// 	app.NotFound(w)
	// 	return
	// }
	//panic("oops")
	s, err := app.snippets.Latest()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// Use the new render helper.
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})

	// for _, snip := range s {
	// 	fmt.Fprintf(w, "%v\n", snip)
	// }

	// data := &templateData{Snippets: s}
	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.ServerError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.ServerError(w, err)
	// }
	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	//using application error log instead of the default log
	// 	//app.errorLog.Println(err.Error())
	// 	app.ServerError(w, err)
	// 	//http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }

	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	//app.errorLog.Println(err.Error())
	// 	//http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	app.ServerError(w, err)
	// }
}
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	//we have to add the ':' otherwise pat will no recognize it
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		//app.errorLog.Println(err.Error())
		//http.Error(w, "Invalid ID", http.StatusNotFound)
		//http.NotFound(w,r)

		app.NotFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotFound(w)

		} else {
			app.ServerError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{Snippet: s})
	//allow rendering of multiple data in template
	// data := &templateData{
	// 	Snippet: s,
	// }
	// files := []string{
	// 	"./ui/html/show.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.ServerError(w, err)
	// 	return
	// }
	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.ServerError(w, err)

	// }
	//fmt.Fprintf(w, "%v", s)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	log.Println("CREATE SNIPPET")
	//pat only accepts POST for this method so no need to check

	// //Allow only POST method to be used per HTTP good practices
	// if r.Method != http.MethodPost {
	// 	//Let user know what Methods are accepted at the endpoint
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	app.ClientError(w, http.StatusMethodNotAllowed)
	// 	//app.errorLog.Println(http.StatusMethodNotAllowed)
	// 	//405
	// 	//http.Error(w, "Method Now allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	// title := "test"
	// content := "test"
	// expires := "7"

	//A new session will be created if the current one is expired or if there
	//wasnt one to begin with. The middleware handles this
	//We are adding data to the "flash" key for the session data
	app.session.Put(r, "flash", "snippet created")

	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}
	// title := r.PostForm.Get("title")
	// content := r.PostForm.Get("content")
	// expires := r.PostForm.Get("expires")

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{
			Form: form,
		})
		return
	}
	//validationErr := make(map[string]string)
	// if strings.TrimSpace(title) == "" {
	// 	validationErr["title"] = "Title cannot be empty"
	// } else if utf8.RuneCountInString(title) > 100 {
	// 	validationErr["title"] = "Title too long"
	// }

	// if strings.TrimSpace(content) == "" {
	// 	validationErr["content"] = "Content cannot be empty"
	// }

	// if strings.TrimSpace(expires) == "" {
	// 	validationErr["expires"] = "Expires cannot be blank"
	// } else if expires != "365" && expires != "7" && expires != "1" {
	// 	validationErr[expires] = "Expires must be 1, 7, or 365"

	// }

	// if len(validationErr) > 0 {
	// 	app.render(w, r, "create.page.tmpl", &templateData{
	// 		FormErrors: validationErr,
	// 		FormData:   r.PostForm,
	// 	})
	// 	return
	// }

	//the values have actaully been anonymously embedded inside the form struct so we
	//can use the form.Get() method to get the valid values
	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.ServerError(w, err)
		return
	}
	//redirect user to the given id
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.MinLength("password", 10)
	form.MatchesPattern("email", forms.EmailRegex)
	form.MaxLength("name", 255)
	form.MaxLength("email", 255)
	form.Required("name", "email", "password")

	//Any errors, redisplay the page with the required fields highlighted
	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.As(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "email already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		} else {
			app.ServerError(w, err)
		}
		return
	}

	app.session.Put(r, "flash", "success")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{Form: forms.New(nil)})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	//grab the information from user, no need to validate since if this matches is should be present
	//in the database and already valid
	form := forms.New(r.PostForm)

	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or Password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.ServerError(w, err)
		}
		return
	}

	// Add the ID of the current user to the session, so that they are now 'logged // in'.
	app.session.Put(r, "authenticatedUserID", id)

	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)

}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	//remove authenticated user id from the current session
	app.session.Remove(r, "authenticatedUserID")

	//flash message successful logout
	app.session.Put(r, "flash", "You've been logged out successfully!")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
