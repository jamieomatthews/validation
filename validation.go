package validation

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/martini-contrib/binding"
)

type Validation struct {
	Errors  binding.Errors
	Request *http.Request
}

type Errors interface {
	WithClass(classification string) Errors
	ForField(name string) Errors
	Get(class, fieldName string) Errors
	Add(fieldNames []string, classification string, message string)
}

type Error interface {
	Fields() []string
	Classification() string
	Message() string
}

//a struct that maps errors.  errors can span multiple feilds,
//and each field can have mutliple errors
type error struct {
	fields         []string // name(s) of the fields involved, if any
	classification string   // error type or category
	message        string   // human-readable or detailed message
}

type errors struct {
	errors []error
}

func (errs *errors) WithClass(classification string) Errors {
	errorsWithClass := &errors{}
	for _, er := range errs.errors {
		if er.Classification() == classification {
			errorsWithClass.errors = append(errorsWithClass.errors, er)
		}
	}
	return errorsWithClass
}

func (errs *errors) ForField(name string) Errors {
	errorsWithField := &errors{}
	for _, er := range errs.errors {
		if stringInSlice(name, er.Fields()) {
			errorsWithField.errors = append(errorsWithField.errors, er)
		}
	}
	return errorsWithField
}

func (errs *errors) Get(class, fieldName string) Errors {
	errToReturn := &errors{}
	for _, er := range errs.errors {
		if stringInSlice(fieldName, er.Fields()) && er.Classification() == class {
			errToReturn.errors = append(errToReturn.errors, er)
		}
	}
	return errToReturn
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (errs *errors) Add(fieldNames []string, classification string, message string) {
	er := error{fields: fieldNames, classification: classification, message: message}
	errs.errors = append(errs.errors, er)
}

func (e *error) Fields() []string {
	return e.fields
}

func (e *error) Classification() string {
	return e.classification
}

func (e *error) Message() string {
	return e.message
}

//represents one 'set' of validation errors.
type Set struct {
	Field          interface{} //a pointer to the passed in feild
	Key            string      //string key pulled from the field
	isValid        bool
	classification string
	message        string
	Validation     *Validation //for now, keep a reference to the validation to map errors back
}

func New(errors binding.Errors, request *http.Request) *Validation {
	return &Validation{Errors: errors, Request: request}
}

func (v *Validation) Validate(field interface{}, key string) *Set {
	s := &Set{Field: field, Key: key, isValid: true, Validation: v}

	return s
}

func (v *Validation) AddError(err binding.Error) {
	v.Errors = append(v.Errors, err)
}

// func (v *Validation) ToJson() []byte {
// 	m := make(map[string][]string)
// 	for err := range v.Errors {
// 		m[err.FieldNames[0]] =
// 	}
// }

// returns true if the validator has 1 or more errors
func (s *Set) HasErrors() bool {
	return len(s.Validation.Errors) > 0
}

func (s *Set) MaxLength(maxLength int) *Set {
	return s.validate(MaxLength{MaxLength: maxLength}, s.Len())
}

func (s *Set) MinLength(minLength int) *Set {
	fmt.Println("Min Length: '", s.Field, "'")
	return s.validate(MinLength{MinLength: minLength}, s.Len())
}

func (s *Set) Match(strMatch string, regex *regexp.Regexp) *Set {
	return s.validate(Matches{Regex: regex}, strMatch)
}

func (s *Set) Email() *Set {
	var emailPattern = regexp.MustCompile("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$")
	return s.validate(Email{Matches{emailPattern}}, s.toString())
}

func (s *Set) Classify(str string) *Set {
	s.classification = str
	return s
}

func (s *Set) Message(message string) *Set {
	s.message = message
	return s
}
func (s *Set) TrimSpace() *Set {
	if str, ok := s.Field.(string); ok {
		s.Field = strings.TrimSpace(str)
	}
	return s
}

func (s *Set) getMessage(val Validator) string {
	if s.message == "" {
		return val.DefaultMessage()
	}
	return s.message
}

func (s *Set) Len() int {
	if str, ok := s.Field.(string); ok {
		return len(str)
	}
	if array, ok := s.Field.([]interface{}); ok {
		return len(array)
	}
	if m, ok := s.Field.(map[interface{}]interface{}); ok {
		return len(m)
	}
	return 0
}

func (s *Set) toString() string {
	if str, ok := s.Field.(string); ok {
		return str
	}
	panic("This method requires a string value")
}

//runs the validation rule, returns true if the rule passed
func (s *Set) validate(validator Validator, obj interface{}) *Set {
	//check if the rule is valid
	if validator.IsValid(obj) {
		fmt.Println("Validated")
		return s
	}

	//else, add a new validation error
	s.Validation.AddError(binding.Error{FieldNames: []string{s.Key}, Classification: s.classification, Message: s.getMessage(validator)})
	s.isValid = false
	return s
}
