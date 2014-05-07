package validation

import (
	"fmt"
	"testing"
)

type User struct {
	name    string `form:name-field`
	age     int
	email   string
	profile string
}

func TestValidation(t *testing.T) {
	user := &User{"John Doe The Fourth", 20, "john@gmail", "John's profile is a long string of text that is more than 20 characters long"}

	errors := errors{}

	v := &Validation{Errors: &errors, Obj: user}

	// //run some validators
	v.Validate(&user.name).MaxLength(15)
	v.Validate(&user.email).Message("Custom Email Validation Message").Classify("email-class").Email()
	v.Validate(&user.profile).TrimSpace().MinLength(10)

	for i := 0; i < v.Errors.Count(); i++ {
		err := v.Errors.At(i)
		fmt.Println("Error: ", err)
	}
}

// func TestReflect(t *testing.T) {
// 	user := &User{"John Doe The Fourth", 20, "john@gmail", "John's profile is a long string of text that is more than 20 characters long"}

// 	attemptReflect(user.name)
// }

// func getStructTag(i interface{}) string{
// 	fmt.Println("Reflect:", reflect.ValueOf(i).Elem().Field(0).
// }
