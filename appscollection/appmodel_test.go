package appscollection

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp("Dixte")
	if app.Name != "Dixte" {
		t.Error("Name wasn't set in the app")
	}
}
