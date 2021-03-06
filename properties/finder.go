package properties

import (
	"github.com/luiz-pv9/dixte/databasemanager"
)

func FindByKey(key string) (*KeyProperties, error) {
	db := databasemanager.Db.Conn
	rows, err := db.Query(`
		SELECT 
			p.property_id, p.key, p.name, p.type, p.is_large, 
			pv.property_values_id, pv.value, pv.count
		FROM properties AS p INNER JOIN property_values AS pv
			ON pv.property_id = p.property_id
		WHERE p.key = $1`, key)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Variables used to scan the values returned from the sql statement
	var (
		property_id        int64
		_key               string
		name               string
		_type              string
		is_large           bool
		property_values_id int64
		value              string
		count              int64
	)
	keyProperties := NewKeyProperties()
	for rows.Next() {
		err = rows.Scan(&property_id, &_key, &name, &_type, &is_large,
			&property_values_id, &value, &count)

		if err != nil {
			return nil, err
		}

		if property := keyProperties.GetProperty(name); property != nil {
			property.AddValue(property_values_id, value, count)
		} else {
			property = keyProperties.AddProperty(property_id, _key, name,
				_type, is_large)
			property.AddValue(property_values_id, value, count)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	rows.Close()
	return keyProperties, nil
}
