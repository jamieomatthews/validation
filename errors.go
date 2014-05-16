package validation

import "fmt"

type Error interface {
	Fields() []string
	Kind() string
	error
}

type Errors interface {
	Add(fields []string, class string, message string)
	Len() int
}

//a struct that maps errors.  errors can span multiple feilds,
//and each field can have mutliple errors
type validationError struct {
	msg    string
	fields []string
	class  string
}

// Satisfy the Error interface for starters
func (e validationError) Error() string    { return e.msg }
func (e validationError) Fields() []string { return e.fields }
func (e validationError) Kind() string     { return e.class }

//a struct that holds an array of pointers to error objects
type ValidationErrors []validationError

// Has determines whether an errors slice has an Error with
// a given classification in it; it does not search on messages
// or field names.
func (e ValidationErrors) Has(class string) bool {
	for _, err := range e {
		if err.Kind() == class {
			return true
		}
	}
	return false
}

// WithClass gets a copy of errors that are classified by the
// the given classification.
func (e ValidationErrors) WithClass(classification string) ValidationErrors {
	var errs ValidationErrors
	for i, err := range e {
		if err.Kind() == classification {
			errs = append(errs, e[i])
		}
	}
	return errs
}

// ForField gets a copy of errors that are associated with the
// field by the given name.
func (e ValidationErrors) ForField(name string) ValidationErrors {
	var errs ValidationErrors
	for _, err := range e {
		for i, fieldName := range err.Fields() {
			if fieldName == name {
				errs = append(errs, e[i])
				break
			}
		}
	}
	return errs
}

// Get gets errors of a particular class for the specified
// field name.
func (e ValidationErrors) Get(class, fieldName string) ValidationErrors {
	var errs ValidationErrors
	for _, err := range e {
		if err.Kind() == class {
			for i, nameOfField := range err.Fields() {
				if nameOfField == fieldName {
					errs = append(errs, e[i])
					break
				}
			}
		}
	}
	return errs
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (e *ValidationErrors) Add(fields []string, class string, message string) {
	fmt.Println("Trying to add an error (myError)")
	err := validationError{
		msg:    message,
		fields: fields,
		class:  class,
	}
	*e = append(*e, err)
}

func (e *ValidationErrors) Len() int {
	return len(*e)
}
