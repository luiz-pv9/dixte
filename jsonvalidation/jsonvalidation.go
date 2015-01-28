package jsonvalidation

import (
	"regexp"
)

type JsonValidator func(val interface{}) bool

// This struct is used to validate objects with unknown structure, such as the
// properties of an event or profile
type JsonKeyValuePairValidator struct {
	Key   JsonValidator
	Value JsonValidator
}

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

// Returns true if the specified argument is a simple json type. A simple json
// value is string, number or boolean.
func AnySimpleValue(val interface{}) bool {
	switch val.(type) {
	case string, float64, bool:
		return true
	default:
		return false
	}
}

// Returns true if the specified argument is a simple json type or null.
// A simple json value is string, number or boolean.
func AnySimpleOrNilValue(val interface{}) bool {
	if val == nil {
		return true
	}
	return AnySimpleValue(val)
}

func CleanObject(properties map[string]interface{},
	validators []*JsonKeyValuePairValidator) map[string]interface{} {
	cleaned := make(map[string]interface{})

	var validator JsonValidator
	for key, val := range properties {
		validator = findValidatorForKey(key, validators)
		if validator != nil && validator(val) == true {
			cleaned[key] = val
		}
	}
	return cleaned
}

func findValidatorForKey(key string,
	validators []*JsonKeyValuePairValidator) JsonValidator {
	for _, validator := range validators {
		if validator.Key(key) == true {
			return validator.Value
		}
	}
	return nil
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

// Generates a function that matches the JsonValidator type to compare a string
// against the specified regexp. The second argument (error) may be returned
// from the regexp compile function, which the user should handle.
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

func OrExpression(rules []JsonValidator) JsonValidator {
	return func(val interface{}) bool {
		for _, validator := range rules {
			if validator(val) == true {
				return true
			}
		}
		return false
	}
}

func AnyArrayByRules(rules []JsonValidator) JsonValidator {
	return func(val interface{}) bool {
		values, ok := val.([]interface{})
		if !ok {
			return false
		}
		matchedCount := 0
		for _, e := range values {
			for _, validator := range rules {
				if validator(e) == true {
					matchedCount += 1
					break
				}
			}
		}
		return matchedCount == len(values)
	}
}
