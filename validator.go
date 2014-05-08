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

type Max struct {
	Max int
}

func (m Max) IsValid(obj interface{}) bool {
	num, ok := obj.(int)
	if ok {
		return num <= m.Max
	}
	return false
}

func (m Max) DefaultMessage() string {
	return fmt.Sprintf("Cannot be larger than %s", m.Max)
}

//min value, only works on integers
type Min struct {
	Min int
}

func (m Min) IsValid(obj interface{}) bool {
	num, ok := obj.(int)
	if ok {
		return num >= m.Min
	}
	return false
}

func (m Min) DefaultMessage() string {
	return fmt.Sprintf("Must be at least %s", m.Min)
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

//matches email (by regex)
type Email struct {
	Matches
}

func (email Email) DefaultMessage() string {
	return fmt.Sprintf("Must be a valid email address")
}
