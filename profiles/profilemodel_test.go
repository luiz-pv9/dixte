package profiles

import (
	"github.com/luiz-pv9/dixte/apps"
	"github.com/luiz-pv9/dixte/databasemanager"
	"github.com/luiz-pv9/dixte/environment"
	"path/filepath"
	"testing"
)

var appToken string

func connectProfileModel() (*databasemanager.Database, *environment.Config) {
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	dc.AssignAppDefaults()
	db, _ := databasemanager.Connect(dc)
	return db, dc
}

func createApp() *apps.App {
	app := apps.NewApp("Dixte")
	app.Token = "Drawing"
	app.Persist()
	return app
}

func clean(db *databasemanager.Database) {
	apps.DeleteAll()
	DeleteAll()
	db.Conn.Close()
}

func profile1(appToken string) *Profile {
	return &Profile{
		ExternalId: "lpvasco",
		AppToken:   appToken,
		Properties: map[string]interface{}{
			"name": "Luiz Paulo",
		},
	}
}

func TestNewProfileTrack(t *testing.T) {
	db, _ := connectProfileModel()
	defer clean(db)
}
