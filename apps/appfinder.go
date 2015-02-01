package apps

import (
	"database/sql"
	"github.com/luiz-pv9/dixte-analytics/databasemanager"
)

func CountRegisteredApps() (int64, error) {
	db := databasemanager.Db.Conn
	row := db.QueryRow("SELECT COUNT(app_id) FROM apps")
	var counter int64
	err := row.Scan(&counter)
	if err != nil {
		return int64(0), err
	}
	return counter, nil
}

// Returns the app found for the specified token or nil if none is found
func FindByToken(token string) (*App, error) {
	db := databasemanager.Db.Conn
	row := db.QueryRow("SELECT app_id, name FROM apps WHERE apps.token = $1",
		token)
	var (
		appId int64
		name  string
	)
	err := row.Scan(&appId, &name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &App{AppId: appId, Name: name, Token: token}, nil
}
