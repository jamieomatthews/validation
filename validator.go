package main

import "fmt"

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
