package databasemanager

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/luiz-pv9/dixte-analytics/environment"
	"io/ioutil"
	"log"
	"path/filepath"
)

var (
	Db *Database
)

type Database struct {
	Conn *sql.DB
}

// Returns the connection to the database using the configuration and
// credentials specified in the Config struct
func Connect(dc *environment.Config) (*Database, error) {
	if Db != nil {
		if err := Db.Conn.Ping(); err == nil {
			log.Printf("RESCUING FROM THE CACHE!\n")
			return Db, nil
		}
	}
	db, err := sql.Open("postgres", dc.Database.ToConnectionArguments())
	if err != nil {
		return nil, err
	}
	dbm := &Database{db}
	Db = dbm // Cache the connection
	return dbm, nil
}

func (db *Database) TablesNames() ([]string, error) {
	data, err := db.Conn.Query(`SELECT table_name FROM information_schema.tables 
		WHERE table_schema = 'public'`)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	tablesNames := make([]string, 0)
	for data.Next() {
		var tableName string
		if err := data.Scan(&tableName); err != nil {
			return nil, err
		}
		tablesNames = append(tablesNames, tableName)
	}
	return tablesNames, nil
}

func (db *Database) Reset() error {
	tablesNames, err := db.TablesNames()
	if err != nil {
		return err
	}
	for _, tableName := range tablesNames {
		_, err := db.Conn.Exec("DROP TABLE " + tableName + " CASCADE")
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Database) CreateMigrationsTable() error {
	command := `
		CREATE SEQUENCE migrations_id_seq;
		CREATE TABLE migrations (
			migration_id INT NOT NULL DEFAULT NEXTVAL('migrations_id_seq'),
			file_name VARCHAR(100) NOT NULL
		);
		ALTER SEQUENCE migrations_id_seq OWNED BY migrations.migration_id;
		ALTER TABLE migrations ADD PRIMARY KEY(migration_id);
		CREATE UNIQUE INDEX file_name_unique_index ON migrations(file_name);
	`
	_, err := db.Conn.Exec(command)
	return err
}

func (db *Database) HasMigrationsTable() bool {
	tablesNames, err := db.TablesNames()
	if err != nil {
		log.Fatal(err)
	}
	for _, tableName := range tablesNames {
		if tableName == "migrations" {
			return true
		}
	}
	return false
}

func (db *Database) resetAndSetupMigrations() error {
	err := db.Reset()
	if err != nil {
		return err
	}
	err = db.CreateMigrationsTable()
	return err
}

func (db *Database) MigratedFiles() ([]string, error) {
	migratedFiles := make([]string, 0)
	rows, err := db.Conn.Query(`SELECT file_name FROM migrations 
		ORDER BY migration_id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var migratedFile string
		if err = rows.Scan(&migratedFile); err != nil {
			return nil, err
		}
		migratedFiles = append(migratedFiles, migratedFile)
	}
	return migratedFiles, nil
}

func (db *Database) loadPendingMigrations(migrationsDir string,
	migratedFiles []string) ([]string, error) {
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return nil, err
	}

	pendingMigrations := make([]string, 0)
	for _, file := range files {
		isMigrated := false
		for _, migrated := range migratedFiles {
			if file.Name() == migrated {
				isMigrated = true
				break
			}
		}
		if !isMigrated {
			pendingMigrations = append(pendingMigrations, file.Name())
		}
	}
	return pendingMigrations, nil
}

func (db *Database) Migrate(migrationsDir string) ([]string, error) {
	if db.HasMigrationsTable() == false {
		err := db.resetAndSetupMigrations()
		if err != nil {
			return nil, err
		}
	}

	migratedFiles, err := db.MigratedFiles()
	if err != nil {
		return nil, err
	}

	pendingMigrations, err := db.loadPendingMigrations(migrationsDir, migratedFiles)
	if err != nil {
		return nil, err
	}

	for _, pendingMigration := range pendingMigrations {
		err = db.migrateFile(filepath.Join(migrationsDir, pendingMigration))
		if err != nil {
			return nil, err
		}
		err = db.storeMigrationRegister(pendingMigration)
		if err != nil {
			return nil, err
		}
	}
	return pendingMigrations, nil
}

func (db *Database) migrateFile(file string) error {
	migration, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	log.Printf("Running migration in file: %v\n", file)
	_, err = db.Conn.Exec(string(migration))
	return err
}

func (db *Database) storeMigrationRegister(migration string) error {
	log.Printf("Storing migration (%v) in the database...\n\n", migration)
	_, err := db.Conn.Exec("INSERT INTO migrations (file_name) VALUES ($1)", migration)
	if err != nil {
		return err
	}
	log.Printf("Migration concluded with success!\n")
	return nil
}
