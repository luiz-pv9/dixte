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
