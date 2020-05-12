package forms

type errors map[string][]string

//Add errors to the given field
func (e errors) Add(key string, message string) {
	e[key] = append(e[key], message)
}

//Get errors from given key
func (e errors) Get(key string) string {
	errs := e[key]
	if len(errs) == 0 {
		return ""
	}

	return errs[0]
}
