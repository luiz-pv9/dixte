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
