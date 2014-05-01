###martini-validate

The idea behind this repo is to give some nice default validation handlers to complement the [Martini Bind Contrib Package](https://github.com/martini-contrib/binding).  

Heres an example of the syntax we are going for: (thanks Matt for suggesting)
```go
func (bp *BlogPost) Validate(errors binding.Errors, req *http.Request) Errors {
    v := validator.New(errors, req)

    v.Validate(&bp.Title).TrimSpace().Length(10, 100).MustNotContain("foo")
    v.Validate(&bp.Subtitle).TrimSpace().MaxLength(50)
    v.Validate(&bp.Content).TrimSpace().MustContain("Go", "Classification", "Custom message")
    v.Validate(&bp.Author.Email).EmailAddress()
    v.Validate(req).HasHeader("X-Foo-Bar") // maybe not useful; but the idea is you can do validation on the request itself

    // then any other custom validation steps you want to perform yourself,
    // more specific to your application or handler usually

    return v.Errors()
}
```

My goal is to make this actually work with any web framework, as it doesnt only apply to Martini.  *For now* I have hard coded in the binding.Errors as the internal error structure, but I will swap this out ASAP for an interface structure that would let anyone utilize it.

Here are some validations I have created so far:
Validations to perform:

-  Minimum Length / Max Length
-  Matches Pattern (takes Regex)
-  Email (uses matches pattern)

As well as some utilities, like
- TrimSpace
- Default (overrides default error)
- Classify (sets classification)


More that I want to add when I have time:

-  EqualTo (other form field)
-  Range length
-  Min/Max value (numbers)
-  Credit Card Number (meets checksum)
-  Credit Card Number (meets checksum)
-  Use matches pattern to do URL, and other pattern like examples




This is **very** work-in-progress, so don't rely on the sytax too much yet. Hope to reach a stable point in a few days.

Ideas inspired from the [jQuery validation plugin](http://jqueryvalidation.org/documentation/) as well as the way .NET MVC handles [model validation](http://www.asp.net/mvc/tutorials/mvc-4/getting-started-with-aspnet-mvc4/adding-validation-to-the-model).

