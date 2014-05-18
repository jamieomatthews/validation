package validation

type Errors interface {
	Add(err Error)
}
type Error interface {
	Fields() []string
	Classification() string
	Message() string
	New(fields []string, classification string, message string) Error //constructor
}

//a struct that maps errors.  errors can span multiple feilds,
//and each field can have mutliple errors
type error struct {
	fields         []string // name(s) of the fields involved, if any
	classification string   // error type or category
	message        string   // human-readable or detailed message
}

//a struct that holds an array of pointers to error objects
type errors []error

func (errs Errors) WithClass(classification string) Errors {
	errorsWithClass := Errors{}
	for _, er := range errs {
		if er.Classification() == classification {
			errorsWithClass = append(errorsWithClass, er)
		}
	}
	return errorsWithClass
}

func (errs Errors) ForField(name string) Errors {
	errorsWithField := Errors{}
	for _, er := range errs {
		if stringInSlice(name, er.Fields()) {
			errorsWithField = append(errorsWithField, er)
		}
	}
	return errorsWithField
}

func (errs Errors) Get(class, fieldName string) Errors {
	errToReturn := Errors{}
	for _, er := range errs {
		if stringInSlice(fieldName, er.Fields()) && er.Classification() == class {
			errToReturn = append(errToReturn, er)
		}
	}
	return errToReturn
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (errs Errors) Len() int {
	return len(errs)
}

func (errs Errors) At(index int) *error {
	return &errs[index]
}

// func (errs Errors) MapErrors() []byte {

// }

func (e *error) New(fields []string, classification string, message string) Error {
	e = &error{fields: fields, classification: classification, message: message}
	return e
}
func (e *error) Fields() []string {
	return e.fields
}

func (e *error) Classification() string {
	return e.classification
}

func (e *error) Message() string {
	return e.message
}
