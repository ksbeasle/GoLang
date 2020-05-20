package main

import (
	"html/template"
	"net/url"
	"path/filepath"
	"time"

	"ksbeasle.net/snippetbox/pkg/forms"
	"ksbeasle.net/snippetbox/pkg/models"
)

type templateData struct {
	Flash           string
	Form            *forms.Form
	CurrentYear     int
	IsAuthenticated bool
	CSRFToken       string
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	FormData        url.Values
	FormErrors      map[string]string
}

func newTemplateCache(dir string) (map[string]*template.Template, error) { // Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}
	// Use the filepath.Glob function to get a slice of all filepaths with
	// the extension '.page.tmpl'. This essentially gives us a slice of all the // 'page' templates for the application.
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}
	// Loop through the pages one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'home.page.tmpl') from the full file path // and assign it to the name variable.
		name := filepath.Base(page)
		// Parse the page template file in to a template set.
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// Use the ParseGlob method to add any 'layout' templates to the // template set (in our case, it's just the 'base' layout at the // moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		// Use the ParseGlob method to add any 'partial' templates to the // template set (in our case, it's just the 'footer' partial at the // moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		// Add the template set to the cache, using the name of the page // (like 'home.page.tmpl') as the key.
		cache[name] = ts
	}
	// Return the map.
	return cache, nil
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}
