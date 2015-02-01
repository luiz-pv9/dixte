package apps

import (
	"github.com/luiz-pv9/dixte-analytics/databasemanager"
	"github.com/luiz-pv9/dixte-analytics/environment"
	"path/filepath"
	"testing"
)

func connectAppModel() (*databasemanager.Database, *environment.Config) {
	dc, _ := environment.LoadConfigFromFile(filepath.Join("..", "config.json"))
	dc.AssignAppDefaults()
	db, _ := databasemanager.Connect(dc)
	return db, dc
}

func TestNewApp(t *testing.T) {
	app := NewApp("Dixte")
	if app == nil {
		t.Error("App wasn't instantiated")
	}

	if app.Name != "Dixte" {
		t.Error("Name wasn't assigned")
	}
}

func TestGenerateToken(t *testing.T) {
	db, dc := connectAppModel()
	defer db.Conn.Close()

	app := NewApp("Dixte")
	app.GenerateToken(dc)
	if app.Token == "" {
		t.Error("Didn't generate token")
	}
	if len(app.Token) != int(dc.App.Token_Size) {
		t.Errorf(`Token (%v) of length (%v) wasn't in the length 
			specified (%v) by the TokenSize`, app.Token, len(app.Token),
			dc.App.Token_Size)
	}
}

func TestGenerateTokenMultipleApps(t *testing.T) {
	db, dc := connectAppModel()
	defer db.Conn.Close()

	app := NewApp("Dixte")
	app.GenerateToken(dc)
	if app.Token == "" {
		t.Error("Didn't generate token")
	}
	if len(app.Token) != int(dc.App.Token_Size) {
		t.Errorf(`Token (%v) of length (%v) wasn't in the length 
			specified (%v) by the TokenSize`, app.Token, len(app.Token),
			dc.App.Token_Size)
	}

	other := NewApp("Other")
	other.GenerateToken(dc)
	if other.Token == "" {
		t.Error("Didn't generate token")
	}
	if other.Token == app.Token {
		t.Errorf("Generated the same token for two apps")
	}
}

func TestCreateApp(t *testing.T) {
	db, _ := connectAppModel()
	defer db.Conn.Close()
	defer DeleteAll()

	app := NewApp("Dixte")
	app.Token = "Francisca"
	err := app.Create()
	if err != nil {
		t.Error(err)
	}

	if app.AppId == int64(0) {
		t.Error("Didn't generate the app_id for the created app")
	}

	count, err := CountRegisteredApps()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("Didn't store the app in the database")
	}
}

func TestCreateAppSameToken(t *testing.T) {
	db, _ := connectAppModel()
	defer db.Conn.Close()
	defer DeleteAll()

	app := NewApp("Dixte")
	app.Token = "Francisca"
	err := app.Create()
	if err != nil {
		t.Error(err)
	}

	other := NewApp("Other")
	other.Token = "Francisca"
	err = other.Create()
	if err == nil {
		t.Error("Didn't raise error when creating the app with same token")
	}

	count, err := CountRegisteredApps()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("Didn't register only the first app")
	}
}

func TestUpdateApp(t *testing.T) {
	db, _ := connectAppModel()
	defer db.Conn.Close()
	defer DeleteAll()

	app := NewApp("Dixte")
	app.Token = "Francisca"
	err := app.Create()
	if err != nil {
		t.Error(err)
	}

	app.Name = "Dixte2"
	app.Token = "Francisca2"
	err = app.Update()
	if err != nil {
		t.Error(err)
	}

	app, err = FindByToken("Francisca2")
	if err != nil {
		t.Error(err)
	}

	if app == nil {
		t.Error("Didn't find app with token Francisca2")
	}

	if app.Name != "Dixte2" {
		t.Error("Didn't update app name")
	}

	app, _ = FindByToken("Francisca")
	if app != nil {
		t.Error("Didn't update record to new token")
	}
}

func TestUpdateAppWithMultipleApps(t *testing.T) {
	db, _ := connectAppModel()
	defer db.Conn.Close()
	defer DeleteAll()

	app1 := NewApp("App1")
	app1.Token = "Token1"
	err := app1.Persist()
	if err != nil {
		t.Error(err)
	}

	app2 := NewApp("App2")
	app2.Token = "Token2"
	err = app2.Persist()
	if err != nil {
		t.Error(err)
	}

	count, err := CountRegisteredApps()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("Didn't register both apps")
	}

	app2.Name = "App3"
	app2.Token = "Token3"
	err = app2.Persist()
	if err != nil {
		t.Error(err)
	}

	count, err = CountRegisteredApps()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("Didn't update second app")
	}

	found1, _ := FindByToken("Token1")
	found2, _ := FindByToken("Token3")
	if found1 == nil || found2 == nil {
		t.Error("Wasn't able to find both apps by their token")
	}
}
