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

func TestTablesNames(t *testing.T) {
	dc, _ := dixteconfig.LoadFromFile(filepath.Join("..", "config.json"))
	db, _ := Connect(dc)
	_, err := db.TablesNames()
	if err != nil {
		t.Error(err)
	}
}

func TestResetDatabase(t *testing.T) {
	dc, _ := dixteconfig.LoadFromFile(filepath.Join("..", "config.json"))
	db, _ := Connect(dc)
	err := db.Reset()
	if err != nil {
		t.Error(err)
	}

	tablesNames, err := db.TablesNames()
	if err != nil {
		t.Error(nil)
	}

	if len(tablesNames) > 0 {
		t.Error("Didn't delete all tables")
	}
}

func TestCreateMigrationsTable(t *testing.T) {
	dc, _ := dixteconfig.LoadFromFile(filepath.Join("..", "config.json"))
	db, _ := Connect(dc)
	err := db.Reset()
	if err != nil {
		t.Error(err)
	}

	err = db.CreateMigrationsTable()
	if err != nil {
		t.Error(err)
	}

	tablesNames, _ := db.TablesNames()
	if len(tablesNames) != 1 || tablesNames[0] != "migrations" {
		t.Error("Didn't create migrations table")
	}
}

func TestMigrationTableExists(t *testing.T) {
	dc, _ := dixteconfig.LoadFromFile(filepath.Join("..", "config.json"))
	db, _ := Connect(dc)
	err := db.Reset()
	if err != nil {
		t.Error(err)
	}

	if db.HasMigrationsTable() != false {
		t.Error("Shouldn't have a migrations table")
	}

	err = db.CreateMigrationsTable()
	if err != nil {
		t.Error(err)
	}

	if db.HasMigrationsTable() != true {
		t.Error("Didn't detect migrations table")
	}
}
