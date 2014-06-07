package validation

import (
	"fmt"
	"testing"
)

func TestValidation(t *testing.T) {
	user := &User{"John Doe The Fourth", 20, "john@gmail", "John's profile is a long string of text that is more than 20 characters long", "34368140066606", "url.com", "", 0}

	//to initialize a validation object, we need to pass in the model object (struct) being validated
	//as well as an array of type Error (interface can be found in errors.go)
	//the point of this is so that you can plug the error interface into your own errors implementation, or use mine

	errors := &ValidationErrors{}
	v := NewValidation(errors, user)

	//you can change the struct tag that is used to map the error key by calling the line bellow
	//feel free to try it out and see how every key except user_name changes, because username is being overridden inline
	//v.KeyTag("custom")

	// //run some validators
	v.Validate(&user.Name).Key("user_name").MaxLength(10)
	v.Validate(&user.Email).Message("Custom Email Validation Message").Classify("email-class").Email()
	v.Validate(&user.Profile).TrimSpace().Range(10, 15)
	v.Validate(&user.CreditCard).CreditCard()
	v.Validate(&user.PageURL).TrimSpace().URL()
	v.Validate(&user.NickName).Required()
	v.Validate(&user.Weight).Required()

	if v.Errors.Len() != 7 {
		t.Errorf("Except 7 errors, but %d catched", v.Errors.Len())
	}
	fmt.Println("Errors len = ", v.Errors.Len())
	//must use type assertion because validate returns an interface
	if myArray, ok := v.Errors.(*ValidationErrors); ok {
		for _, err := range *myArray {
			fmt.Printf("\t%s: '%s')\n", err.fields, err.msg)
		}
	}

	//fmt.Println("\n\n", v.MapErrors())
}
