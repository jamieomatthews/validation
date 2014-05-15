package validation

import (
	"fmt"
	"testing"
)

type User struct {
	Name    string `form:"aName"`
	Age     int    `form:"age"`
	Email   string `form:"email"`
	Profile string `form:"profile"`
}

func TestValidation(t *testing.T) {
	user := &User{"John Doe The Fourth", 20, "john@gmail", "John's profile is a long string of text that is more than 20 characters long"}

	//to initialize a validation object, we need to pass in the model object (struct) being validated
	//as well as an array of type Error (interface can be found in errors.go)
	//the point of this is so that you can plug the error interface into your own errors implementation, or use mine

	errors := []error{}
	v := &Validation{Errors: errors, Obj: user}

	// //run some validators
	v.Validate(&user.Name).Key("user_name").MaxLength(10)
	v.Validate(&user.Email).Message("Custom Email Validation Message").Classify("email-class").Email()
	v.Validate(&user.Profile).TrimSpace().MaxLength(10)

	fmt.Println("Errors len = ", len(v.Errors))
	for _, err := range v.Errors {
		fmt.Println("Error: ", err)
		//fmt.Println(err.fields, err.Message)
	}
}
