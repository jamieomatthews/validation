package validation

import (
	"fmt"
	"testing"
)

type User struct {
	name    string
	age     int
	email   string
	profile string
}

func TestValidation(t *testing.T) {
	user := &User{"John Doe The Fourth", 20, "john@gmail.com", "John's profile is a long string of text that is more than 20 characters long"}

	v := NewDefault()

	// //run some validators
	v.Validate(user.name, "name").MaxLength(15)
	v.Validate(user.email, "email").Message("Custom Email Validation Message").Classify("email-class").Email()
	v.Validate(user.profile, "profile").TrimSpace().MinLength(10)

	for i := 0; i < v.Errors.Count(); i++ {
		err := v.Errors.At(i)
		fmt.Println("Error: ", err)
	}
}
