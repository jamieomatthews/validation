package validation

import (
	"fmt"
	"reflect"
	"regexp"
	"time"
)

//each validator will just implement this interface,
//which will be enough for the validation object to validate and create errors
type Validator interface {
	IsValid(interface{}) bool
	DefaultMessage() string
}

type Required struct {
	Key string
}

func (r Required) IsValid(obj interface{}) bool {
	if obj == nil {
		return false
	}

	if str, ok := obj.(string); ok {
		return len(str) > 0
	}
	if b, ok := obj.(bool); ok {
		return b
	}
	if i, ok := obj.(int); ok {
		return i != 0
	}
	if t, ok := obj.(time.Time); ok {
		return !t.IsZero()
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() > 0
	}
	return true
}

func (r Required) DefaultMessage() string {
	return fmt.Sprint("This Field Is Required")
}

//max length validator
type MaxLength struct {
	MaxLength int
}

func (max MaxLength) DefaultMessage() string {
	return fmt.Sprintf("Maximum Length is %d", max.MaxLength)
}

func (max MaxLength) IsValid(obj interface{}) bool {
	num, ok := obj.(int)
	if ok {
		return num <= max.MaxLength
	}
	return false
}

//min length validator
type MinLength struct {
	MinLength int
}

func (min MinLength) DefaultMessage() string {
	return fmt.Sprintf("Minimum Length is %d", min.MinLength)
}

func (min MinLength) IsValid(obj interface{}) bool {
	num, ok := obj.(int)
	if ok {
		return num >= min.MinLength
	}
	return false
}

// Range valid within Min, Max inclusive.
type Range struct {
	MinLength
	MaxLength
}

func (r Range) IsValid(obj interface{}) bool {
	return r.MinLength.IsValid(obj) && r.MaxLength.IsValid(obj)
}

func (r Range) DefaultMessage() string {
	return fmt.Sprintf("Length must be within %d and %d", r.MinLength.MinLength, r.MaxLength.MaxLength)
}

//matches pattern validator
//primarily meant for internal use for other validators, like email or credit card
type Matches struct {
	Regex *regexp.Regexp
}

func (m Matches) IsValid(obj interface{}) bool {
	str := obj.(string)
	return m.Regex.MatchString(str)
}

func (m Matches) DefaultMessage() string {
	return fmt.Sprintf("Must match %s", m.Regex)
}

//short hand for not matching
type NoMatch struct {
	Matches
}

func (m NoMatch) IsValid(obj interface{}) bool {
	return !m.Matches.IsValid(obj)
}

func (m NoMatch) DefaultMessage() string {
	return fmt.Sprintf("Must not match %s", m.Matches.Regex)
}

//very simple (not fool proof) email pattern
var emailPattern = regexp.MustCompile("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$")

//matches email (by regex)
type Email struct {
	Matches
}

func (email Email) DefaultMessage() string {
	return fmt.Sprintf("Must be a valid email address")
}

//Matches Visa, MasterCard, American Express, Diners Club, Discover, and JCB cards
//Note that this in no way validates the actual card, just that it could be a valid card
var creditCardPattern = regexp.MustCompile("^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$")

type CreditCard struct {
	Matches
}

func (credit CreditCard) DefaultMessage() string {
	return fmt.Sprintf("Must be a valid credit card number")
}

//matches most urls
//source: http://blog.mattheworiordan.com/post/13174566389/url-regular-expression-for-links-with-or-without-the
var urlPattern = regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[\-;:&=\+\$,\w]+@)?[A-Za-z0-9\.\-]+|(?:www\.|[\-;:&=\+\$,\w]+@)[A-Za-z0-9\.\-]+)((?:\/[\+~%\/\.\w\-_]*)?\??(?:[\-\+=&;%@\.\w_]*)#?(?:[\.\!\/\\\w]*))?)`)

type URL struct {
	Matches
}

func (url URL) DefaultMessage() string {
	return fmt.Sprintf("Must be a valid URL")
}
