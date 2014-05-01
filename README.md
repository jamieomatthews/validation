###martini-validate

The idea behind this repo is to give some nice default validation handlers to complement the [Martini Bind Contrib Package](https://github.com/martini-contrib/binding).  

Heres an example of the syntax we are going for: (thanks [Matt](https://github.com/mholt) for suggesting)
```go

func (contactRequest ContactRequest) Validate(errors binding.Errors, req *http.Request) binding.Errors {

	v := validation.New(errors, req)

	// //run some validators
	v.Validate(contactRequest.FullName, "full_name").MaxLength(20)
	v.Validate(contactRequest.Email, "email").Default("Custom Email Validation Message").Classify("email-class").Email()
	v.Validate(contactRequest.Comments, "comments").TrimSpace().MinLength(10)

	return v.Errors
}

type ContactRequest struct {
	FullName string `form:"full_name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Subject  string `form:"subject"`
	Comments string `form:"comments"`
}
```

My goal is to make this actually work with any web framework, as it doesnt only apply to Martini.  *For now* I have hard coded in the binding.Errors as the internal error structure, but I will swap this out ASAP for an interface structure that would let anyone utilize it.

Here are some validations I have created so far:
Validations to perform:

-  **MaxLength(maxLength int) / MinLength(minLength int)** - works for strings, arrays, and maps
-  **Matches(regex *regexp.Regexp)** - returns true if it meets the regex
-  **Email()** - uses matches pattern

As well as some utilities, like

-   **TrimSpace()** - trims whitespace
-   **Default(message string)** - overrides default error message
-   **Classify(classification string)** - sets classification


More that I want to add when I have time:

-  EqualTo (other form field)
-  Range length
-  Min/Max value (numbers)
-  Credit Card Number (meets checksum)
-  Credit Card Number (meets checksum)
-  Use matches pattern to do URL, and other pattern like examples




This is **very** work-in-progress, so don't rely on the sytax too much yet. Hope to reach a stable point in a few days.

Ideas inspired from the [jQuery validation plugin](http://jqueryvalidation.org/documentation/) as well as the way .NET MVC handles [model validation](http://www.asp.net/mvc/tutorials/mvc-4/getting-started-with-aspnet-mvc4/adding-validation-to-the-model).

