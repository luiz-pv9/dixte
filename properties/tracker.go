package properties

// This struct is just a list of properties with some helper functions included
type KeyProperties struct {
	Properties []*Property
}

func (kp *KeyProperties) GetProperty(name string) *Property {
	for _, property := range kp.Properties {
		if property.Name == name {
			return property
		}
	}
	return nil
}

func (kp *KeyProperties) AddProperty(id int64, key, name, _type string) *Property {
	property := NewProperty(id, key, name, _type)
	kp.Properties = append(kp.Properties, property)
	return property
}

func (kp *KeyProperties) GetTotalCount() int64 {
	var total int64 = 0
	for _, property := range kp.Properties {
		total += property.GetTotalCount()
	}
	return total
}

type Property struct {
	PropertyId int64
	Key        string
	Name       string
	Type       string
	Values     []*Value
}

func (p *Property) AddValue(id int64, value string, count int64) *Value {
	_value := NewValue(id, p.PropertyId, value, count)
	p.Values = append(p.Values, _value)
	return _value
}

func (p *Property) GetValue(value string) *Value {
	for _, _value := range p.Values {
		if _value.Value == value {
			return _value
		}
	}
	return nil
}

func (p *Property) GetTotalCount() int64 {
	var total int64 = 0
	for _, value := range p.Values {
		total += value.Count
	}
	return total
}

type Value struct {
	PropertyValueId int64
	PropertyId      int64
	Value           string
	Count           int64
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

func Track(key string, properties map[string]interface{}) error {
	return nil
}
