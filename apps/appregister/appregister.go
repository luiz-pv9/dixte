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

func generateRandomToken(size int) string {
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(rb)
}

func (app *App) GenerateToken(dc *dixteconfig.DixteConfig) error {
	searching := true
	for searching {
		token := generateRandomToken(int(dc.App.Token_Size))
		appFound, err := appfinder.ByToken(token)
		if appFound == nil {
			app.Token = token
			searching = false
		}
	}
}
