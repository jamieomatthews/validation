package validation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type Validation struct {
	Errors  Errors
	Obj     interface{} //the top-most model struct being validated
	Request *http.Request
}

//represents one 'set' of validation errors.
type Set struct {
	Field          interface{} //a pointer to the passed in feild
	key            string      //string key pulled from the field
	isValid        bool
	classification string
	message        string
	Error          validationError
	Validation     *Validation //for now, keep a reference to the validation to map errors back
}

func New(errors Errors, obj interface{}) *Validation {
	return &Validation{Errors: errors, Obj: obj}
}

func (v *Validation) Validate(field interface{}) *Set {

	s := &Set{Field: field, isValid: true, Validation: v}
	key := v.getKeyForField(field)
	s.key = key

	return s
}

func (v *Validation) MapErrors() string {
	if errSlice, ok := v.Errors.(*ValidationErrors); ok {
		if b, err := json.Marshal(errSlice); err == nil {
			return string(b)
		}
	}
	return ""
}

//experimenting with trying to match the field with the passed in struct
func (v *Validation) getKeyForField(passedField interface{}) string {
	typObj := reflect.TypeOf(v.Obj)
	valObj := reflect.ValueOf(v.Obj)

	typField := reflect.TypeOf(passedField)
	valField := reflect.ValueOf(passedField)

	//if our struct is a pointer, dereference it
	if typObj.Kind() == reflect.Ptr {
		typObj = typObj.Elem()
		valObj = valObj.Elem()
	}

	//if our passed in field is a pointer, dereference it
	if typField.Kind() == reflect.Ptr {
		typField = typField.Elem()
		valField = valField.Elem()
	}

	for i := 0; i < typObj.NumField(); i++ {
		field := typObj.Field(i)
		fieldValue := valObj.Field(i).Interface()
		passedValue := valField.Interface()
		if passedValue == fieldValue {
			return field.Tag.Get("form")
		}
	}
	return "not found"
}

func (s *Set) Key(key string) *Set {
	s.key = key
	return s
}
func (s *Set) Required() *Set {
	return s.validate(Required{}, s.Field)
}

// returns true if the validator has 1 or more errors
func (s *Set) HasErrors() bool {
	return s.Validation.Errors.Len() > 0
}

func (s *Set) MaxLength(maxLength int) *Set {
	return s.validate(MaxLength{MaxLength: maxLength}, s.Len())
}

func (s *Set) MinLength(minLength int) *Set {
	fmt.Println("Min Length: '", s.Field, "'")
	return s.validate(MinLength{MinLength: minLength}, s.Len())
}

func (s *Set) Max(max int) *Set {
	return s.validate(Max{max}, s.Field)
}

func (s *Set) Min(min int) *Set {
	return s.validate(Min{min}, s.Field)
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

	switch x := s.Field.(type) {
	case string:
		return len(x)

	case *string:
		return len(*x)

	case []interface{}:
		return len(x)
	case *[]interface{}:
		return len(*x)
	case map[interface{}]interface{}:
		return len(x)
	case *map[interface{}]interface{}:
		return len(*x)
	}

	return 0
}

func (s *Set) toString() string {
	if str, ok := s.Field.(string); ok {
		return str
	}

	if str, ok := s.Field.(*string); ok {
		return *str
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

	fmt.Println("Not Validated, adding error")
	//else, add a new validation error
	s.Validation.Errors.Add([]string{s.key}, s.classification, s.getMessage(validator))
	//s.Validation.Errors = append(s.Validation.Errors, er)
	s.isValid = false
	return s
}
