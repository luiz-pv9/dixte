package appfinder

import (
	"github.com/luiz-pv9/dixte-analytics/databasemanager"
	"github.com/luiz-pv9/dixte-analytics/dixteconfig"
	"testing"
)

func connection() *databasemanager.Database {
	dc, _ := dixteconfig.LoadFromFile(filepath.Join("..", "config.json"))
	dixteconfig.LoadFromFile(filepath)
	db, _ := databasemanager.Connect(dc)
	return db
}

func TestByToken(t *testing.T) {
	db := connection()
	app, err := ByToken("any-token", db.Conn)
	if err != nil {
		t.Error(err)
	}

	if app != nil {
		t.Error("Found an app the shouldn't exist.")
	}
}
