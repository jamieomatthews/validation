package validation

type Validator interface {
	IsValid(interface{}) bool
	DefaultMessage() string
}
