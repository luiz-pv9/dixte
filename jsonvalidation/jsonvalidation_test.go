package jsonvalidation

import (
	"encoding/json"
	"testing"
)

func TestAnyString(t *testing.T) {
	var fn JsonValidator
	fn = AnyString
	if fn("foobar") != true {
		t.Error("Didn't match foobar string")
	}

	if fn("") != true {
		t.Error("Didn't match empty string")
	}

	if fn(nil) != false {
		t.Error("Matched nil argument")
	}

	if fn(10) != false {
		t.Error("Matched number argument")
	}
}

func TestAnyNumber(t *testing.T) {
	var fn JsonValidator
	fn = AnyNumber

	if fn(float64(30)) != true {
		t.Error("Didn't match 30")
	}

	if fn(float64(31.12312312)) != true {
		t.Error("Didn't match with decimal points")
	}

	if fn(nil) != false {
		t.Error("Matched a nil value")
	}

	if fn("123") != false {
		t.Error("Matched a string value")
	}
}

func TestAnyBoolean(t *testing.T) {
	var fn JsonValidator
	fn = AnyBoolean

	if fn(false) != true {
		t.Error("Didn't match false value")
	}

	if fn(true) != true {
		t.Error("Didn't match true value")
	}

	if fn(1) != false {
		t.Error("Matched number value as boolean")
	}

	if fn(nil) != false {
		t.Error("Matched a nil value")
	}
}

func TestAnySimpleValue(t *testing.T) {
	var fn JsonValidator
	fn = AnySimpleValue

	if fn(false) != true || fn(true) != true {
		t.Error("Didn't match boolean values")
	}

	if fn(float64(123.23)) != true || fn(float64(0)) != true {
		t.Error("Didn't match json numeric value")
	}

	if fn(nil) != false {
		t.Error("Matched nil value")
	}

	if fn("foobar") != true || fn("") != true {
		t.Error("Didn't match string value")
	}

	strLists := []string{"a", "b", "c"}
	if fn(strLists) != false {
		t.Error("Matched an array list")
	}
}

func TestAnySimpleOrNilValue(t *testing.T) {
	var fn JsonValidator
	fn = AnySimpleOrNilValue

	if fn(nil) != true {
		t.Error("Didn't match nil value (and it should)")
	}

	if fn("foobar") != true || fn(float64(123)) != true || fn(true) != true {
		t.Error("Didn't match simple values not nil")
	}
}

func TestExactString(t *testing.T) {
	var fn JsonValidator
	fn = ExactString("luiz")
	if fn("luiz") != true {
		t.Error("Didn't match exact string (luiz)")
	}

	if fn("Luiz") != false {
		t.Error("Matched string with different case (Luiz)")
	}

	if fn("") != false {
		t.Error("Matched empty string")
	}

	if fn(nil) != false {
		t.Error("Matched nil argument")
	}

	fn = ExactString("123")
	if fn(123) != false {
		t.Error("Matched number argument although equal('123', 123)")
	}
}

func TestRegexpString(t *testing.T) {
	var fn JsonValidator
	fn, err := RegexpString("[ab]foo")
	if err != nil {
		t.Error("Didn't compile specified regexp")
	}

	if fn("bfoo") != true {
		t.Error("Didn't match with the specified regex")
	}

	if fn("afoo") != true {
		t.Error("Didn't match with the specified regex")
	}

	if fn("cfoo") != false {
		t.Error("Matched value not in regexp")
	}

	if fn("") != false {
		t.Error("Matched empty string")
	}

	if fn(nil) != false {
		t.Error("Matched nil argument")
	}

	if fn(123) != false {
		t.Error("Matched numeric argument")
	}

	fn, err = RegexpString("asd#${!ÇLAÇSD>>.xz.,!@##!$%[")
	if err == nil {
		t.Error("Didn't raise error with invalid regexp")
	}
}

func TestCleanObject(t *testing.T) {
	data := `
	{
		"name": "Luiz",
		"age": 20,
		"premium": false,
		"$push.colors": ["red"],
		"numbers": [10, 20]
	}`
	var properties map[string]interface{}
	json.Unmarshal([]byte(data), &properties)

	// One rule with any string -> any simple value
	var rules []*JsonKeyValuePairValidator
	rules = []*JsonKeyValuePairValidator{
		&JsonKeyValuePairValidator{
			AnyString, AnySimpleOrNilValue,
		},
	}

	cleaned := CleanObject(properties, rules)
	if cleaned["name"] != "Luiz" || cleaned["age"] != float64(20) || cleaned["premium"] != false {
		t.Error("Removed simple values from the hash that shouldn't be removed.")
	}

	if cleaned["$push.colors"] != nil || cleaned["numbers"] != nil {
		t.Error("Didn't remove properties that were not specified in the rules.")
	}

	// Multiple rules with regexp matching a value before the any string
	// This should give the same output as the previous one
	premiumAge, _ := RegexpString("premium|age")
	rules = []*JsonKeyValuePairValidator{
		&JsonKeyValuePairValidator{
			premiumAge, AnySimpleOrNilValue,
		},
		&JsonKeyValuePairValidator{
			AnyString, AnyString,
		},
	}

	cleaned = CleanObject(properties, rules)
	if cleaned["name"] != "Luiz" || cleaned["age"] != float64(20) || cleaned["premium"] != false {
		t.Error("Removed simple values from the hash that shouldn't be removed.")
	}
	if cleaned["$push.colors"] != nil || cleaned["numbers"] != nil {
		t.Error("Didn't remove properties that were not specified in the rules.")
	}
}
