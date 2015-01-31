package databasemanager

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/luiz-pv9/dixte-analytics/dixteconfig"
	"io/ioutil"
	"path/filepath"
)

type Migration struct {
	file string
}

func Connect(dc *dixteconfig.DixteConfig) (*sql.DB, error) {
	config := make(map[string]string)
}

// TODO
// The migrations should actually be stored in the database.
// Create a table, store the file name and run only what is not there yet.
func RunMigrations(dc *dixteconfig.DixteConfig) error {
	migrationsPath := filepath.Join("..", "migrations")
	files := ioutil.ReadDir(migrationsPath)
	for _, file := range files {
	}
}
