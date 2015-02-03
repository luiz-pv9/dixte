package properties

import (
	"github.com/luiz-pv9/dixte-analytics/databasemanager"
	"github.com/luiz-pv9/dixte-analytics/environment"
	"path/filepath"
	"testing"
)

var appToken string

func connectProfileModel() (*databasemanager.Database, *environment.Config) {
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	dc.AssignAppDefaults()
	db, _ := databasemanager.Connect(dc)
	return db, dc
}

func clean(db *databasemanager.Database) {
	DeleteAll()
	db.Conn.Close()
}

func TestTracking(t *testing.T) {
	db, _ := connectProfileModel()
	defer clean(db)
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
}

func TestTrackingTypeDetection(t *testing.T) {
	db, _ := connectProfileModel()
	defer clean(db)
}

func TestTrackingIncrementingValue(t *testing.T) {
	db, _ := connectProfileModel()
	defer clean(db)
}

func TestTrackingMultipleValuesForSameProperty(t *testing.T) {
	db, _ := connectProfileModel()
	defer clean(db)
	// "Luiz" and "Paulo" for same property: "name"
}

func TestTrackingLargeCollections(t *testing.T) {
	db, _ := connectProfileModel()
	defer clean(db)
}
