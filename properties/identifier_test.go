package properties

import (
	"testing"
)

func TestNumberIdentify(t *testing.T) {
	var val interface{}
	val = int64(18)
	if IdentifyType(val) != "number" {
		t.Error("Didn't identify int64 value")
	}

	val = float64(123.456789)
	if IdentifyType(val) != "number" {
		t.Error("Didn't identify float64 value")
	}
}

func TestBooleanIdentify(t *testing.T) {
	var val interface{}
	val = false
	if IdentifyType(val) != "boolean" {
		t.Error("Didn't identify false boolean type")
	}

	val = true
	if IdentifyType(val) != "boolean" {
		t.Error("Didn't identify true boolean type")
	}
}

func TestArrayIdentify(t *testing.T) {
	var val interface{}
	val = []interface{}{"str", "str2"}

	if IdentifyType(val) != "array" {
		t.Error("Didn't identify array type")
	}

	val = []interface{}{"str", 20}
	if IdentifyType(val) != "array" {
		t.Error("Didn't identify array with multiple types")
	}

}

func TestObjectIdentify(t *testing.T) {
	var val interface{}
	val = map[string]interface{}{
		"name": "Luiz",
		"age":  20,
	}

	if IdentifyType(val) != "object" {
		t.Error("Didn't identify object")
	}
}

func TestNullIdentify(t *testing.T) {
	var val interface{}
	val = nil
	if IdentifyType(val) != "null" {
		t.Error("Didn't identify null value")
	}
}

func TestStringIdentify(t *testing.T) {
	var val interface{}
	val = "what"
	if IdentifyType(val) != "string" {
		t.Error("Didn't identify string value")
	}

	val = "123"
	if IdentifyType(val) != "string" {
		t.Error("Didn't identify string value")
	}
}
