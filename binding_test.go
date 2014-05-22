package validation

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/martini-contrib/binding"
)

func TestBinding(t *testing.T) {
	user := &User{Name: "John Doe The Fourth", Age: 20, Email: "john@gmail", Profile: "John's profile is a long string of text that is more than 20 characters long"}

	//to initialize a validation object, we need to pass in the model object (struct) being validated
	//as well as an array of type Error (interface can be found in errors.go)
	//in this case, we're going to use the martini binding package errors

	errors := &binding.Errors{}
	v := NewValidation(errors, user)

	// //run some validators
	v.Validate(&user.Name).Key("user_name").MaxLength(10)
	v.Validate(&user.Email).Message("Custom Email Validation Message").Classify("email-class").Email()
	v.Validate(&user.Profile).TrimSpace().MaxLength(10)

	//must use type assertion because validate returns an interface
	if errs, ok := v.Errors.(*binding.Errors); ok {
		errOutput, _ := json.Marshal(errs)
		fmt.Println("=====================\n\n", string(errOutput))
	}

}
