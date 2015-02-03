package properties

// This struct is just a list of properties with some helper functions included
// to manipulate all properties of a given key. A key will be usually the
// app token concatenated to an event type.
type KeyProperties struct {
	Properties []*Property
}

// A property is equivalent of the key of a map of properties. For example, in
// the following json `{ "name": "Luiz" }` the property would be "name" and
// "Luiz" one of it's values. Type detection is also stored in the property
// table, so in the same example since "Luiz" is a string, the type of the
// property would be "string".
//
// The key is used to identify who owns this property, and will usually be
// the app token concatenated with an event type.
type Property struct {
	PropertyId int64
	Key        string
	Name       string
	Type       string
	Values     []*Value
}

// A property have multiple values, and each value has a count associated with
// it. In the following json `{"browser": "Firefox"}` a property with the name
// "browser" would be created and the value "Firefox" would be stored with the
// count set to 1. If more properties with value "Firefox" are tracked the
// count will increment. When a value is removed, the counter is decremented.
type Value struct {
	PropertyValueId int64
	PropertyId      int64
	Value           string
	Count           int64
}

// Finds the property with the specified name and returns it. Nil is returned
// if nothing is found
func (kp *KeyProperties) GetProperty(name string) *Property {
	if kp.Properties == nil {
		return nil
	}
	for _, property := range kp.Properties {
		if property.Name == name {
			return property
		}
	}
	return nil
}

func (kp *KeyProperties) AddProperty(id int64, key, name, _type string) *Property {
	property := NewProperty(id, key, name, _type)
	if kp.Properties == nil {
		kp.Properties = make([]*Property, 0)
	}
	kp.Properties = append(kp.Properties, property)
	return property
}

func (kp *KeyProperties) GetTotalCount() int64 {
	var total int64 = 0
	if kp.Properties == nil {
		return total
	}
	for _, property := range kp.Properties {
		total += property.GetTotalCount()
	}
	return total
}

func (p *Property) AddValue(id int64, value string, count int64) *Value {
	_value := NewValue(id, p.PropertyId, value, count)
	if p.Values == nil {
		p.Values = make([]*Value, 0)
	}
	p.Values = append(p.Values, _value)
	return _value
}

func (p *Property) GetValue(value string) *Value {
	if p.Values == nil {
		return nil
	}
	for _, _value := range p.Values {
		if _value.Value == value {
			return _value
		}
	}
	return nil
}

func (p *Property) GetTotalCount() int64 {
	var total int64 = 0
	if p.Values == nil {
		return total
	}
	for _, value := range p.Values {
		total += value.Count
	}
	return total
}

func NewKeyProperties() *KeyProperties {
	keyProperties := &KeyProperties{}
	keyProperties.Properties = make([]*Property, 0)
	return keyProperties
}

func NewProperty(id int64, key, name, _type string) *Property {
	property := &Property{
		PropertyId: id,
		Key:        key,
		Name:       name,
		Type:       _type,
		Values:     make([]*Value, 0),
	}
	return property
}

func NewValue(id, propertyId int64, value string, count int64) *Value {
	_value := &Value{
		PropertyValueId: id,
		PropertyId:      propertyId,
		Value:           value,
		Count:           count,
	}
	return _value
}
