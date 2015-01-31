package databasemanager

import (
	"github.com/luiz-pv9/dixte-analytics/dixteconfig"
	"io/ioutil"
	"path/filepath"
)

// TODO
// The migrations should actually be stored in the database.
// Create a table, store the file name and run only what is not there yet.
func RunMigrations(dc *dixteconfig.DixteConfig) error {
	files := ioutil.ReadDir(filepath.Join("..", "migrations"))
}
