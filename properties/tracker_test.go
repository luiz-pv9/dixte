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
	// err := Track("foobar", map[string]interface{}{
	// 	"name": "Luiz Paulo",
	// })
	// if err != nil {
	// 	t.Error(err)
	// }
}
