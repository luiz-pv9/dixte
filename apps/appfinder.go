package apps

import (
	"database/sql"
	"github.com/luiz-pv9/dixte-analytics/databasemanager"
)

// Returns the app found for the specified token or nil if none is found
func FindByToken(token string) (*App, error) {
	db := databasemanager.Db.Conn
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
	return &App{Id: id, Name: name, Token: token}, nil
}
