package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

//method to initialize a new Form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Implement a Required method to check that specific fields in the form
// data are present and not blank. If any fields fail this check, add the
//and add appropriate message to the form errors
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This cannot be empty")
		}
	}
}

//Method to check max length
func (f *Form) MaxLength(field string, length int) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) > length {
		f.Errors.Add(field, fmt.Sprintf("field is longer than %d", length))
	}
}

//check the specific field in form matches
//at least one of the specific opts
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "Invalid")
}

//boolean method that will return true if there were no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// This returns a *regexp.Regexp object, or panics in the event of an error.
//^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$
var EmailRegex = regexp.MustCompile("^\\w+@[a-zA-Z_]+?\\.[a-zA-Z]{2,3}$")

//Check min length of user
func (f *Form) MinLength(field string, length int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < length {
		f.Errors.Add(field, fmt.Sprintf("need length greater than %d", length))
	}
}

func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if pattern.MatchString(value) == false {
		f.Errors.Add(field, "Invalid pattern")
	}
}
