package appregister

import (
	"database/sql"
	"github.com/luiz-pv9/dixte-analytics/apps/appmodel"
)

func Persist(app *appmodel.App, db *sql.DB) error {
	if app.Id == 0 {
		Create(app, db)
	} else {
		Update(app, db)
	}
}

func Create(app *appmodel.App, db *sql.DB) error {
}

func Update(app *appmodel.App, db *sql.DB) error {
}
