package appconfig

import (
	"os"
	"testing"
)

func currentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

func TestLoadFromFile(t *testing.T) {
	curDir := currentDirectory()
	appConfig, err := LoadFromFile(curDir + "/01_data.json")
	if err != nil {
		t.Error(err)
	}
	if appConfig.Port != float64(8080) {
		t.Error("Didn't read the specified port of 8080")
	}

	if appConfig.Database.Host != "localhost" {
		t.Error("Didn't read the specified host of the database")
	}
}
