package validation

type ValidationError struct {
	Key, Value string
}

type Validation struct {
	Errors []*ValidationError
}

// returns true if the validator has 1 or more errors
func (v *Validation) HasErrors() bool {
	return len(v.Errors) > 0
}
