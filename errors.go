package validation

type Errors interface {
	WithClass(classification string) Errors
	ForField(name string) Errors
	Get(class, fieldName string) Errors
	Add(fieldNames []string, classification string, message string)

	//for iterating
	Count() int
	At(index int) Error
}

type Error interface {
	Fields() []string
	Classification() string
	Message() string
}

//a struct that maps errors.  errors can span multiple feilds,
//and each field can have mutliple errors
type error struct {
	fields         []string // name(s) of the fields involved, if any
	classification string   // error type or category
	message        string   // human-readable or detailed message
}

//a struct that holds an array of pointers to error objects
type errors struct {
	errors []*error
}

func (errs *errors) WithClass(classification string) Errors {
	errorsWithClass := &errors{}
	for _, er := range errs.errors {
		if er.Classification() == classification {
			errorsWithClass.errors = append(errorsWithClass.errors, er)
		}
	}
	return errorsWithClass
}

func (errs *errors) ForField(name string) Errors {
	errorsWithField := &errors{}
	for _, er := range errs.errors {
		if stringInSlice(name, er.Fields()) {
			errorsWithField.errors = append(errorsWithField.errors, er)
		}
	}
	return errorsWithField
}

func (errs *errors) Get(class, fieldName string) Errors {
	errToReturn := &errors{}
	for _, er := range errs.errors {
		if stringInSlice(fieldName, er.Fields()) && er.Classification() == class {
			errToReturn.errors = append(errToReturn.errors, er)
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

func (errs *errors) Add(fieldNames []string, classification string, message string) {
	er := error{fields: fieldNames, classification: classification, message: message}
	errs.errors = append(errs.errors, &er)
}

func (errs *errors) Count() int {
	return len(errs.errors)
}

func (errs *errors) At(index int) Error {
	return errs.errors[index]
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
