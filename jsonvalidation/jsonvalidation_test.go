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

func TestAnyArrayByRules(t *testing.T) {
	var fn JsonValidator
	rules := []JsonValidator{AnyString}
	fn = AnyArrayByRules(rules)

	onlyStrings := []interface{}{"foobar"}
	if fn(onlyStrings) != true {
		t.Error("Didn't match array of strings")
	}

	stringsAndNumbers := []interface{}{"foobar", 20}
	if fn(stringsAndNumbers) != false {
		t.Error("Matched array with number in the middle")
	}

	emptyArray := []interface{}{}
	if fn(emptyArray) != true {
		t.Error("Did not match empty array")
	}

	multipleRules := []JsonValidator{AnyString, AnyBoolean}
	fn = AnyArrayByRules(multipleRules)
	stringsAndBooleans := []interface{}{"foobar", true, false, "what"}
	if fn(stringsAndBooleans) != true {
		t.Error("Didn't match array of only strings and booleans")
	}

	stringsAndBooleansAndNumbers := []interface{}{"foobar", true, 10, "foo"}
	if fn(stringsAndBooleansAndNumbers) != false {
		t.Error("Matched numeric value when only strings and booleans are allowed")
	}

	stringsAndBooleansWithNil := []interface{}{"foobar", true, nil, "was"}
	if fn(stringsAndBooleansWithNil) != false {
		t.Error("Matched array with nil value as an element")
	}
}

func TestOrExpression(t *testing.T) {
	var fn JsonValidator
	fn = OrExpression([]JsonValidator{AnyString, AnyNumber})

	if fn("foo") != true {
		t.Error("Didn't match string value")
	}

	if fn(float64(123)) != true {
		t.Error("Didn't match a number value")
	}

	if fn(true) != false {
		t.Error("Matched a boolean value")
	}

	if fn(nil) != false {
		t.Error("Matched a nil value")
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

	// Multiple rules with push allowed for array of strings or a single string
	pushRegex, _ := RegexpString("\\$push\\.[^\\.]*")
	rules = []*JsonKeyValuePairValidator{
		&JsonKeyValuePairValidator{
			pushRegex, AnyArrayByRules([]JsonValidator{AnyString}),
		},
		&JsonKeyValuePairValidator{
			AnyString, AnySimpleValue,
		},
	}
	cleaned = CleanObject(properties, rules)
	if cleaned["name"] != "Luiz" || cleaned["age"] != float64(20) || cleaned["premium"] != false {
		t.Error("Removed simple values from the hash that shouldn't be removed.")
	}
	if cleaned["$push.colors"] == nil {
		t.Error("Didn't match $push operator with array of strings")
	}
	if cleaned["numbers"] != nil {
		t.Error("Matched array of numbers that should not be matched")
	}
}

func TestNestedObjectMatching(t *testing.T) {
	matcher := AnyObjectByRules([]*JsonKeyValuePairValidator{
		&JsonKeyValuePairValidator{
			ExactString("name"), AnyObjectByRules([]*JsonKeyValuePairValidator{
				&JsonKeyValuePairValidator{
					ExactString("first"), AnyString,
				},
				&JsonKeyValuePairValidator{
					ExactString("last"), AnyString,
				},
			}),
		},
		&JsonKeyValuePairValidator{
			ExactString("age"), AnyNumber,
		},
	})

	obj := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Luiz",
			"last":  "Vasconcellos",
		},
		"age": float64(20),
	}

	if matcher(obj) != true {
		t.Error("Didn't match object: %v", obj)
	}

	if matcher(nil) != false {
		t.Error("Matched nil argument")
	}

	obj = map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Luiz",
		},
		"age": float64(20),
	}

	if matcher(obj) != true {
		t.Error("Didn't match object: %v (missing fields)", obj)
	}

	obj = map[string]interface{}{
		"name": map[string]interface{}{
			"first":  "Luiz",
			"last":   "Vasconcellos",
			"middle": "Wait for it",
		},
		"age": float64(20),
	}

	if matcher(obj) != false {
		t.Error("Matched object with unauthorized fields: %v", obj)
	}
}
