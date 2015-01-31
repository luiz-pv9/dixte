package databasemanager

import (
	"github.com/luiz-pv9/dixte-analytics/dixteconfig"
	"path/filepath"
	"testing"
)

func TestConnection(t *testing.T) {
	dc, err := dixteconfig.LoadFromFile(filepath.Join("..", "config.json"))
	if err != nil {
		t.Error(err)
	}
	dc.AssignDefaults()
	db, err := Connect(dc)
	if err != nil {
		t.Error(err)
	}
	err = db.Conn.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestMigrationTableExists(t *testing.T) {
}