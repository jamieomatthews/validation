###martini-validate

The idea behind this package is to give some nice default validation handlers for input handlers.  

The package was developed with respect to the [Martini Binding](https://github.com/martini-contrib/binding) repo, but can actually be used very generically by implementing the Errors interface.

###Usage

```go
func (contactRequest ContactRequest) Validate(errors binding.Errors, req *http.Request) binding.Errors {

    v := NewValidation(&errors, contactRequest)

	// //run some validators
	v.Validate(&contactRequest.FullName).Key("fullname").MaxLength(20)
	v.Validate(&contactRequest.Email).Default("Custom Email Validation Message").Classify("email-class").Email()
	v.Validate(&contactRequest.Comments).TrimSpace().MinLength(10)

	return *v.Errors.(*binding.Errors)
}

type ContactRequest struct {
	FullName string `form:"full_name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Subject  string `form:"subject"`
	Comments string `form:"comments"`
}
```

This will generate something like:
```json
[
    {
        "fieldNames": [
            "fullname"
        ],
        "message": "Maximum Length is 20"
    },
    {
        "fieldNames": [
            "email"
        ],
        "classification": "email-class",
        "message": "Custom Email Validation Message"
    },
    {
        "fieldNames": [
            "comments"
        ],
        "message": "Minimum Length is 10"
    }
]
```

###Configuration

By default, the validator will grab the `form` key out of the struct tag to use as the output key.  This is nice because if you're using the form tag already you don't have to write out any additional keys, which keeps things DRY.

To change what struct tag will be used to map the errors, use `Validation.KeyTag(string)`.

Keys can also be specified on per validation basis by chaining `.Key(string)`.  Note that you must use this *before* you call the validator, as errors get mapped immeditaly after you call a validator. 

Also, make sure to pass the struct fields in as pointers if you want the validator to be able to make changes to the underlying values.  For example, `TrimSpace()` cant actually trim the space unless it recieves a pointer.

###API
Pre-Build Validators:

-  **MaxLength(maxLength int) / MinLength(minLength int)** - works on strings, arrays, and maps
-  **Range(min, max)** - short hand for calling MinLength(int) & MaxLength(int)
-  **Matches(regex *regexp.Regexp)** - returns true if it meets the regex
-  **NoMatch(regex *regexp.Regex)** 
-  **Email()** - uses matches pattern
-  **CreditCard()** matches Visa, MasterCard, American Express, Diners Club, Discover, and JCB cards
-  **URL()** matches most url schemes

As well as some utilities, like

-   **TrimSpace()** - trims whitespace
-   **Message(message string)** - overrides default error message
-   **Classify(classification string)** - sets classification
-   **Key(str string)** - set the key to map out to

More that I want to add when I have time:

-  **EqualTo** (other form field)
-  Use matches pattern to do other pattern like examples


###Contributions welcome!
**Todo's:**

- Write some validators on the http.Request.  For example, HasHeader(), etc
- Improve the syntax to handle multi-field errors
- Add some of the validators listed as want-to-haves above
- **Improve Test Coverage**

Ideas inspired from the [jQuery validation plugin](http://jqueryvalidation.org/documentation/) as well as the way .NET MVC handles [model validation](http://www.asp.net/mvc/tutorials/mvc-4/getting-started-with-aspnet-mvc4/adding-validation-to-the-model).

