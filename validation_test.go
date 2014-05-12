package validation

import (
	"fmt"
	"testing"
)

type User struct {
	Name    string `form:"aName"`
	Age     int
	Email   string
	Profile string
}

func TestValidation(t *testing.T) {
	user := &User{"John Doe The Fourth", 20, "john@gmail", "John's profile is a long string of text that is more than 20 characters long"}

	errors := errors{}

	v := &Validation{Errors: &errors, Obj: user}

	// //run some validators
	v.Validate(&user.Name).MaxLength(10)
	//v.Validate(&user.Email).Message("Custom Email Validation Message").Classify("email-class").Email()
	//v.Validate(&user.Profile).TrimSpace().MinLength(10)

	for i := 0; i < v.Errors.Len(); i++ {
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
