package properties

import (
	"github.com/luiz-pv9/dixte/databasemanager"
	"github.com/luiz-pv9/dixte/environment"
	"path/filepath"
	"testing"
)

func connectTracker() (*databasemanager.Database, *environment.Config) {
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	dc.AssignAppDefaults()
	db, _ := databasemanager.Connect(dc)
	return db, dc
}

func cleanTracker(db *databasemanager.Database) {
	DeleteAll()
	db.Conn.Close()
}

func TestTracking(t *testing.T) {
	db, _ := connectTracker()
	defer cleanTracker(db)
	err := Track("foobar", map[string]interface{}{
		"name": "Luiz Paulo",
	})
	if err != nil {
		t.Error(err)
	}

	var count int64
	row := db.Conn.QueryRow("SELECT COUNT(*) FROM properties")
	err = row.Scan(&count)

	if err != nil {
		t.Error(err)
	}

	if count != int64(1) {
		t.Error("Didn't register the properties in the database")
	}

	// Inserted values
	kp, err := FindByKey("foobar")
	if err != nil {
		t.Error(err)
	}
	name := kp.GetProperty("name")
	if name == nil {
		t.Error("Didn't register name proprety in the database")
	}

	if name.Type != "string" || name.IsLarge != false {
		t.Error("Didn't insert correct attributes of property")
	}

	if len(name.Values) != 1 {
		t.Error("Didn't allocate value for name")
	}

	if name.GetValue("Luiz Paulo") == nil {
		t.Error("Didn't store value for name")
	}

	if name.GetValue("Luiz Paulo").Count != int64(1) {
		t.Error("Didn't update count for value")
	}

	if kp.GetTotalCount() != int64(1) {
		t.Error("Increment values more than it should")
	}
}

func TestTrackingTypeDetection(t *testing.T) {
	db, _ := connectTracker()
	defer cleanTracker(db)
}

func TestTrackingIncrementingValue(t *testing.T) {
	db, _ := connectTracker()
	defer cleanTracker(db)
}

func TestTrackingMultipleValuesForSameProperty(t *testing.T) {
	db, _ := connectTracker()
	defer cleanTracker(db)
	// "Luiz" and "Paulo" for same property: "name"
}

func TestTrackingLargeCollections(t *testing.T) {
	db, _ := connectTracker()
	defer cleanTracker(db)
}

func TestTrackingConcurrent(t *testing.T) {
	db, _ := connectTracker()
	defer cleanTracker(db)
}
