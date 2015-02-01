package apps

import (
	"github.com/luiz-pv9/dixte-analytics/databasemanager"
	"github.com/luiz-pv9/dixte-analytics/environment"
	"path/filepath"
	"testing"
)

func connectAppFinder() *databasemanager.Database {
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	db, _ := databasemanager.Connect(dc)
	return db
}

func TestByTokenNotFound(t *testing.T) {
	db := connectAppFinder()
	defer db.Conn.Close()
	app, err := FindByToken("random-token")

	if err != nil {
		t.Error(err)
	}

	if app != nil {
		t.Error("Found an app the shouldn't exist.")
	}
}

func TestByTokenFound(t *testing.T) {
	db := connectAppFinder()
	defer db.Conn.Close()
}
