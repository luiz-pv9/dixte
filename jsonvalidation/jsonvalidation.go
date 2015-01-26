package jsonvalidation

import (
	"regexp"
)

type JsonValidator func(val interface{}) bool

// Returns true if the specified argument is of the type string
func AnyString(val interface{}) bool {
	_, ok := val.(string)
	return ok
}

// The AnyNumber JsonValidator only checks for numeric values in the format
// of float64 because it's the format go native json will unmarshal any number.
func AnyNumber(val interface{}) bool {
	_, ok := val.(float64)
	return ok
}

// Returns true if the specified argument is of the type boolean (true or false)
func AnyBoolean(val interface{}) bool {
	_, ok := val.(bool)
	return ok
}

// Generates a function that matches the JsonValidator type to compare string
// against the specified value. Usually will be used for keys in a JSON object.
func ExactString(str string) JsonValidator {
	return func(val interface{}) bool {
		stringVal, ok := val.(string)
		if ok == false {
			return false
		}
		return str == stringVal
	}
}

func RegexpString(expression string) (JsonValidator, error) {
	regex, err := regexp.Compile(expression)
	if err != nil {
		return nil, err
	}
	validator := func(val interface{}) bool {
		stringVal, ok := val.(string)
		if !ok {
			return false
		}
		return regex.MatchString(stringVal)
	}
	return validator, nil
}
