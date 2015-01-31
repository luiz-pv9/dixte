package databasemanager

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/luiz-pv9/dixte-analytics/dixteconfig"
	"log"
)

type Database struct {
	Conn *sql.DB
}

// Returns the connection to the database using the configuration and
// credentials specified in the DixteConfig struct
func Connect(dc *dixteconfig.DixteConfig) (*Database, error) {
	db, err := sql.Open("postgres", dc.Database.ToConnectionArguments())
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
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
