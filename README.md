###martini-validate

The idea behind this repo is to give some nice default validation handlers to complement the [Martini Bind Contrib Package](https://github.com/martini-contrib/binding).  

I have mocked up some really common ones, but please send a Pull Request if you develop other general ones you want to share.

Here are some validations I have created so far:
Validations to perform:

1.  Minimum Length / Max Length
2.  Matches Pattern (takes Regex)


More that I want to add when I have time:

-  EqualTo (other form field)
-  Range length
-  Min/Max value (numbers)
-  Credit Card Number (meets checksum)
-  Credit Card Number (meets checksum)
-  Use matches pattern to do URL, and other pattern like examples


For anyone taking a look, I still need to finish this up a bit, as well as re-write the error mapping method to comply with the new [error struct](https://github.com/martini-contrib/binding/issues/22).  Bare with me, this will be done soon. I'll then slowly start chipping away at the second list of validators

Ideas inspired from the [jQuery validation plugin](http://jqueryvalidation.org/documentation/) as well as the way .NET MVC handles [model validation](http://www.asp.net/mvc/tutorials/mvc-4/getting-started-with-aspnet-mvc4/adding-validation-to-the-model).

