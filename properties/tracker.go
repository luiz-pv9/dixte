package properties

import (
	"github.com/luiz-pv9/dixte/databasemanager"
	"log"
	"sync"
)

var (
	muUpdate *sync.Mutex = &sync.Mutex{}
)

// The Track method is only allowed to do updates. This is because the update
// lock will happen in the database, so there is no need to sync updating
// counters. But when creating the records there is chance of conflict, so the
// creation is synced here in the app using, and the track must acquire the right
// to perform the update.
func Track(key string, properties map[string]interface{}) error {
	// err := acquireTrack(key, properties)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func acquireTrack(key string, properties map[string]interface{}) error {
	muUpdate.Lock()
	defer muUpdate.Unlock()

	keyProperties, err := FindByKey(key)
	if err != nil {
		return err
	}

	for name, val := range properties {
		property := keyProperties.GetProperty(name)
		if property == nil {
			propertyId, err := createProperty(key, name, IdentifyType(val),
				false)
			if err != nil {
				return err
			}
			propertyValueId, err := createPropertyValue(propertyId, val, int64(0))
			if err != nil {
				return err
			}
			log.Println(propertyValueId)
			continue
		}
		value := property.GetValue(ToStoreValue(val))
		if value == nil {
			// Run create value SQL
		}
	}

	return nil
}

func createProperty(key, name, _type string, isLarge bool) (int64, error) {
	db := databasemanager.Db.Conn
	var propertyId int64
	err := db.QueryRow(`INSERT INTO properties (key, name, type, is_large) 
		VALUES ($1, $2, $3, $4) RETURNING property_id`,
		key, name, _type, isLarge).Scan(&propertyId)
	return propertyId, err
}

func createPropertyValue(propertyId int64, val interface{},
	count int64) (int64, error) {
	db := databasemanager.Db.Conn

	var propertyValueId int64
	err := db.QueryRow(`INSERT INTO property_values (property_id, value, count)
		VALUES ($1, $2, $3) RETURNING property_value_id`,
		propertyId, ToStoreValue(val), count).Scan(&propertyValueId)

	return propertyValueId, err
}

func ToStoreValue(val interface{}) string {
	return val.(string)
}
