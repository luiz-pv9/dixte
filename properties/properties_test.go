package properties

import (
	"testing"
)

func TestKeyPropertiesAllocation(t *testing.T) {
	kp := NewKeyProperties()
	if kp == nil {
		t.Errorf("Didn't allocate KeyProperties")
	}

	property := kp.AddProperty(int64(10), "foobar", "name", "string", false)

	if len(kp.Properties) != 1 {
		t.Errorf("Didn't add property to array of properties: %v", kp)
	}

	if property != kp.Properties[0] {
		t.Error("Didn't return the created property from the KeyProperties")
	}

	if property.PropertyId != int64(10) {
		t.Error("Didn't assign property id to property")
	}

	if property.Key != "foobar" {
		t.Error("Didn't assign key to property")
	}

	if property.Name != "name" {
		t.Error("Didn't assign name to property")
	}

	if property.Type != "string" {
		t.Error("Didn't assign type to property")
	}
}

func TestValuePropertiesAllocation(t *testing.T) {
	property := NewProperty(int64(10), "foobar", "name", "string", false)
	if property == nil {
		t.Error("Didn't allocate the property")
	}

	if len(property.Values) != 0 {
		t.Error("Didn't allocate values array when creating the property")
	}

	value := property.AddValue(int64(20), "Luiz", int64(2))
	if value == nil {
		t.Error("Didn't allocate value")
	}

	if len(property.Values) != 1 || value != property.Values[0] {
		t.Error("Didn't add the value to the array in the properties")
	}

	if value.PropertyValueId != int64(20) {
		t.Error("Didn't assign id to property value")
	}

	if value.Value != "Luiz" {
		t.Error("Didn't assign value to property value")
	}

	if value.Count != int64(2) {
		t.Error("Didn't assign the count to the property value")
	}
}

func TestGetValue(t *testing.T) {
	property := NewProperty(int64(10), "foobar", "name", "string", false)
	if property.GetValue("foo") != nil {
		t.Error("Didn't return nil when searching for a non existing value")
	}

	val := property.AddValue(int64(1), "Luiz", int64(3))
	if property.GetValue("Luiz") != val {
		t.Error("Didn't recover property value")
	}

	if property.GetValue("luiz") != nil {
		t.Error("Case-sensitive search for value in property search")
	}
}

func TestGetTotalCountOfProperty(t *testing.T) {
	property := NewProperty(int64(10), "foobar", "name", "string", false)

	if property.GetTotalCount() != int64(0) {
		t.Error("Didn't sum count of all properties: %v", property.GetTotalCount())
	}

	property.AddValue(int64(1), "Luiz", int64(2))

	if property.GetTotalCount() != int64(2) {
		t.Error("Didn't sum count of all properties: %v", property.GetTotalCount())
	}

	property.AddValue(int64(2), "Paulo", int64(3))

	if property.GetTotalCount() != int64(5) {
		t.Error("Didn't sum count of all properties: %v", property.GetTotalCount())
	}
}

func TestGetTotalCountOfKeyProperties(t *testing.T) {
	kp := NewKeyProperties()

	if kp.GetTotalCount() != int64(0) {
		t.Error("Total count wasn't zero")
	}

	property := kp.AddProperty(int64(1), "foobar", "age", "number", false)

	if kp.GetTotalCount() != int64(0) {
		t.Error("Total count wasn't zero")
	}

	property.AddValue(int64(1), "20", int64(2))

	if kp.GetTotalCount() != int64(2) {
		t.Error("Total count wasn't 2")
	}

	property.AddValue(int64(2), "21", int64(3))

	if kp.GetTotalCount() != int64(5) {
		t.Error("Total count wasn't 5")
	}

	otherProperty := kp.AddProperty(int64(2), "foobar", "name", "string", false)
	otherProperty.AddValue(int64(3), "Luiz", int64(3))

	if kp.GetTotalCount() != int64(8) {
		t.Error("Total count wasn't 8")
	}
}
