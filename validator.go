package validation

import (
	"fmt"
	"regexp"
)

//each validator will just implement this interface,
//which will be enough for the validation object to validate and create errors
type Validator interface {
	IsValid(interface{}) bool
	DefaultMessage() string
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
