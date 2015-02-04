package properties

import (
	"database/sql"
	"github.com/luiz-pv9/dixte/databasemanager"
	"github.com/luiz-pv9/dixte/environment"
	"path/filepath"
	"testing"
)

func connectFinder() (*databasemanager.Database, *environment.Config) {
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	dc.AssignAppDefaults()
	db, _ := databasemanager.Connect(dc)
	return db, dc
}

func cleanFinder(db *databasemanager.Database) {
	DeleteAll()
	db.Conn.Close()
}

func trackProperties1(key string, db *sql.DB) error {
	var propertyId int64
	err := db.QueryRow(`INSERT INTO properties (key, name, type, is_large)
		VALUES ($1, $2, $3, $4) RETURNING property_id`,
		key, "name", "string", false).Scan(&propertyId)

	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO property_values (property_id, value, count)
		VALUES ($1, $2, $3) RETURNING property_values_id`,
		propertyId, "Luiz", int64(2))

	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO property_values (property_id, value, count)
		VALUES ($1, $2, $3) RETURNING property_values_id`,
		propertyId, "Paulo", int64(3))

	return err
}

func trackProperties2(key string, db *sql.DB) error {
	var propertyId int64
	err := db.QueryRow(`INSERT INTO properties (key, name, type, is_large)
		VALUES ($1, $2, $3, $4) RETURNING property_id`,
		key, "name", "string", false).Scan(&propertyId)

	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO property_values (property_id, value, count)
		VALUES ($1, $2, $3) RETURNING property_values_id`,
		propertyId, "Luiz", int64(2))

	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO property_values (property_id, value, count)
		VALUES ($1, $2, $3) RETURNING property_values_id`,
		propertyId, "Paulo", int64(3))

	err = db.QueryRow(`INSERT INTO properties (key, name, type, is_large)
		VALUES ($1, $2, $3, $4) RETURNING property_id`,
		key, "age", "number", false).Scan(&propertyId)

	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO property_values (property_id, value, count)
		VALUES ($1, $2, $3) RETURNING property_values_id`,
		propertyId, "18", int64(1))

	return err
}

func trackProperties3(key string, db *sql.DB) error {
	var propertyId int64
	err := db.QueryRow(`INSERT INTO properties (key, name, type, is_large)
		VALUES ($1, $2, $3, $4) RETURNING property_id`,
		key, "name", "string", false).Scan(&propertyId)

	err = db.QueryRow(`INSERT INTO properties (key, name, type, is_large)
		VALUES ($1, $2, $3, $4) RETURNING property_id`,
		key, "age", "number", false).Scan(&propertyId)

	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO property_values (property_id, value, count)
		VALUES ($1, $2, $3) RETURNING property_values_id`,
		propertyId, "18", int64(1))

	return err
}

func TestFindByKey(t *testing.T) {
	db, _ := connectFinder()
	defer cleanFinder(db)

	err := trackProperties1("foobar", db.Conn)
	if err != nil {
		t.Error(err)
	}

	kp, err := FindByKey("foobar")
	if err != nil {
		t.Error(err)
	}

	if kp == nil {
		t.Error("Didn't alocate key properties")
	}

	if len(kp.Properties) != 1 {
		t.Error("Didn't alocate properties array")
	}

	nameProp := kp.GetProperty("name")

	if nameProp == nil {
		t.Error("Didn't load name property")
	}

	if nameProp.IsLarge != false {
		t.Error("Didn't load is_large property of name")
	}

	if nameProp.Name != "name" {
		t.Error("Didn't load name property of name")
	}

	if nameProp.Type != "string" {
		t.Error("Didn't load type property of name")
	}

	if nameProp.PropertyId == int64(0) {
		t.Error("Didn't load property id of name")
	}

	if len(nameProp.Values) != 2 {
		t.Error("Didn't allocate array of values for name")
	}

	luiz := nameProp.GetValue("Luiz")
	paulo := nameProp.GetValue("Paulo")
	if luiz == nil || paulo == nil {
		t.Error("Didn't find values for name")
	}

	if luiz.Count != int64(2) || paulo.Count != int64(3) {
		t.Error("Didn't load count for values of name")
	}
}

func TestFindByKeyWithMultipleProperties(t *testing.T) {
	db, _ := connectFinder()
	defer cleanFinder(db)

	// Registers 2 properties: "name" and "age"
	// "name" will have two values: "Luiz" and "Paulo"
	// "age" will have one value: "18"
	err := trackProperties2("foobar", db.Conn)

	if err != nil {
		t.Error(err)
	}

	kp, err := FindByKey("foobar")

	if err != nil {
		t.Error(err)
	}

	if len(kp.Properties) != 2 {
		t.Error("Didn't load name and age properties")
	}

	name := kp.GetProperty("name")
	age := kp.GetProperty("age")

	if name == nil || age == nil {
		t.Error("Didn't find name or age of properties")
	}

	if len(age.Values) != 1 || len(name.Values) != 2 {
		t.Error("Didn't load values for name or age")
	}

	if age.GetValue("18") == nil {
		t.Error("Didn't load value for age")
	}
}

// This test just makes sure the SQL is returning just properties
// related to a specific key
func TestFindByKeyWithMultipleKeys(t *testing.T) {
	db, _ := connectFinder()
	defer cleanFinder(db)

	err := trackProperties1("foobar1", db.Conn)
	if err != nil {
		t.Error(err)
	}
	err = trackProperties3("foobar2", db.Conn)

	if err != nil {
		t.Error(err)
	}

	kp, err := FindByKey("foobar1")

	if kp == nil {
		t.Error("Didn't load property")
	}

	if kp.GetProperty("name") == nil {
		t.Error("Didn't load name property")
	}

	if kp.GetProperty("age") != nil {
		t.Error("Load property of another key")
	}
}
