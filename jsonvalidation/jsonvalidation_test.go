package jsonvalidation

import "testing"

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
}
