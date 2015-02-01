package databasemanager

import (
	"github.com/luiz-pv9/dixte-analytics/environment"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestConnection(t *testing.T) {
	dc, err := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	if err != nil {
		t.Error(err)
	}
	dc.AssignDefaults()
	db, err := Connect(dc)
	if err != nil {
		t.Error(err)
	}
	defer db.Conn.Close()
	err = db.Conn.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestTablesNames(t *testing.T) {
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	db, _ := Connect(dc)
	defer db.Conn.Close()
	_, err := db.TablesNames()
	if err != nil {
		t.Error(err)
	}
}

func TestResetDatabase(t *testing.T) {
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	db, _ := Connect(dc)
	defer db.Conn.Close()
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
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	db, _ := Connect(dc)
	defer db.Conn.Close()
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
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	db, _ := Connect(dc)
	defer db.Conn.Close()
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

func TestMigrate(t *testing.T) {
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	db, _ := Connect(dc)
	defer db.Conn.Close()
	db.Reset()
	migratedFiles, err := db.Migrate(filepath.Join("..", "migrations"))
	if err != nil {
		t.Error(err)
	}
	if len(migratedFiles) < 1 {
		t.Error("Didn't run any migration")
	}
	tables, _ := db.TablesNames()
	if len(tables) < 1 {
		t.Error("Didn't create any tables")
	}
	files, _ := ioutil.ReadDir(filepath.Join("..", "migrations"))
	latestMigratedFiles, _ := db.MigratedFiles()
	if len(files) != len(latestMigratedFiles) {
		t.Error("Didn't store the migrations in the database")
	}

	// Running migrations again shouldn't change anything
	migratedFiles, err = db.Migrate(filepath.Join("..", "migrations"))
	if err != nil {
		t.Error(err)
	}
	if len(migratedFiles) != 0 {
		t.Error("Migrations should not be runned")
	}
}

func TestMigrateNewFiles(t *testing.T) {
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	db, _ := Connect(dc)
	defer db.Conn.Close()
	db.Reset()
	migratedFiles, err := db.Migrate(filepath.Join("..", "migrations"))
	if err != nil {
		t.Error(err)
	}
	log.Print("Migrated files")
	log.Print(migratedFiles)
	if len(migratedFiles) < 1 {
		t.Error("Didn't migrate any files")
	}

	migration := "SELECT schema_name FROM information_schema.schemata"
	newMigrationPath := filepath.Join("..", "migrations", "999-test.sql")
	err = ioutil.WriteFile(newMigrationPath, []byte(migration), 0644)
	if err != nil {
		t.Error(err)
	}

	previousMigratedFiles, _ := db.MigratedFiles()
	previousMigratedFilesLength := len(previousMigratedFiles)
	migratedFiles, err = db.Migrate(filepath.Join("..", "migrations"))
	if err != nil {
		t.Error(err)
	}

	if len(migratedFiles) != 1 {
		t.Error("Didn't pick up new migration file")
	}

	currentMigratedFiles, _ := db.MigratedFiles()
	if len(currentMigratedFiles) != previousMigratedFilesLength+1 {
		t.Error("Didnt include the new migration in the database table")
	}

	err = os.Remove(newMigrationPath)
	if err != nil {
		t.Error(err)
	}
}
