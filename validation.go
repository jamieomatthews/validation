package validation

import (
	"regexp"

	"github.com/martini-contrib/binding"
)

type Validation struct {
	Errors binding.Errors
}

// returns true if the validator has 1 or more errors
func (v *Validation) HasErrors() bool {
	return len(v.Errors.Fields) > 0
}

func (v *Validation) MaxLength(n int, maxLength int, fieldName string) bool {
	return v.validate(MaxLength{MaxLength: maxLength}, n, fieldName)
}

func (v *Validation) MinLength(n int, minLength int, fieldName string) bool {
	return v.validate(MinLength{MinLength: minLength}, n, fieldName)
}

func (v *Validation) Match(strMatch string, regex *regexp.Regexp, fieldName string) bool {
	return v.validate(Matches{Regex: regex}, strMatch, fieldName)
}

func (v *Validation) Email(email string, fieldName string) bool {
	var emailPattern = regexp.MustCompile("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$")
	return v.validate(Email{Matches{emailPattern}}, email, fieldName)
}

//runs the validation rule, returns true if the rule passed
func (v *Validation) validate(validator Validator, obj interface{}, fieldName string) bool {
	//check if the rule is valid
	if validator.IsValid(obj) {
		return true
	}

	//else, add a new validation error
	v.Errors.Fields[fieldName] = validator.DefaultMessage()
	return false
}
