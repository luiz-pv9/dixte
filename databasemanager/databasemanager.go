package databasemanager

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/luiz-pv9/dixte-analytics/dixteconfig"
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

func (db *Database) HasMigrationsTable() bool {
	return false
}
