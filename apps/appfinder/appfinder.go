package appfinder

import (
	"database/sql"
	"github.com/luiz-pv9/dixte-analytics/apps/appmodel"
)

// Returns the app found for the specified token or nil if none is found
func ByToken(token string, db *sql.DB) (*appmodel.App, error) {
	row := db.QueryRow("SELECT app_id, name FROM apps WHERE apps.token = $1",
		token)
	var (
		id   int
		name string
	)
	err := row.Scan(&id, &name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &appmodel.App{Id: id, Name: name, Token: token}, nil
}
