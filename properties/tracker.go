package properties

import (
	"github.com/luiz-pv9/dixte/databasemanager"
	"sync"
)

var (
	// Mutex used to syncrhonize access to properties
	// used by acquireTrack
	muUpdate *sync.Mutex = &sync.Mutex{}
)

// The Track method is only allowed to do updates. This is because the update
// lock will happen in the database, so there is no need to sync updating
// counters. But when creating the records there is chance of conflict, so the
// creation is synced here in the app using, and the track must acquire the right
// to perform the update.
func Track(key string, properties map[string]interface{}) error {
	kp, err := acquireTrack(key, properties)
	if err != nil {
		return err
	}
	db := databasemanager.Db.Conn

	for name, val := range properties {
		property := kp.GetProperty(name)
		value := property.GetValue(ToStoreValue(val))
		db.Exec(`UPDATE property_values SET count = count + 1
			WHERE property_values_id = $1`, value.PropertyValueId)
	}
	return nil
}

// This function uses the muUpdate mutex to syncrhonize access to properties.
// The Track function is only allowed to perform updates incrementing the
// counter, and this function makes sure the register exists when performing
// the update (the lock happens in the database).
// The mutex is necessary to prevent conflict when only creating the properties.
//
// The mutex used by acquireTrack will eventually be replaced by a distributed
// mutex managed by redis.
func acquireTrack(key string,
	properties map[string]interface{}) (*KeyProperties, error) {
	muUpdate.Lock()
	defer muUpdate.Unlock()

	keyProperties, err := FindByKey(key)
	if err != nil {
		return nil, err
	}

	var propertyId, propertyValueId int64
	for name, val := range properties {
		property := keyProperties.GetProperty(name)
		if property == nil {
			var (
				_type   string = IdentifyType(val)
				isLarge bool   = false
				count   int64  = int64(0)
				value   string = ToStoreValue(val)
			)

			propertyId, err = createProperty(key, name, _type, isLarge)
			if err != nil {
				return nil, err
			}

			property = keyProperties.AddProperty(propertyId, key, name, _type,
				isLarge)
			propertyValueId, err = createPropertyValue(propertyId, value, count)
			if err != nil {
				return nil, err
			}
			property.AddValue(propertyValueId, value, count)
			continue
		}

		value := ToStoreValue(val)
		memValue := property.GetValue(value)
		if memValue == nil {
			var count int64 = int64(0)
			propertyValueId, err = createPropertyValue(property.PropertyId,
				value, count)
			if err != nil {
				return nil, err
			}
			property.AddValue(propertyValueId, value, count)
		}
	}
	return keyProperties, nil
}

func createProperty(key, name, _type string, isLarge bool) (int64, error) {
	db := databasemanager.Db.Conn
	var propertyId int64
	err := db.QueryRow(`INSERT INTO properties (key, name, type, is_large) 
		VALUES ($1, $2, $3, $4) RETURNING property_id`,
		key, name, _type, isLarge).Scan(&propertyId)
	return propertyId, err
}

func createPropertyValue(propertyId int64, val string,
	count int64) (int64, error) {
	db := databasemanager.Db.Conn

	var propertyValueId int64
	err := db.QueryRow(`INSERT INTO property_values (property_id, value, count)
		VALUES ($1, $2, $3) RETURNING property_values_id`,
		propertyId, val, count).Scan(&propertyValueId)

	return propertyValueId, err
}

func ToStoreValue(val interface{}) string {
	return val.(string)
}
