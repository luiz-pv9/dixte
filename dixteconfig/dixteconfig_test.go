package dixteconfig

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
	dixteConfig, err := LoadFromFile(curDir + "/01_data.json")
	if err != nil {
		t.Error(err)
	}

	if dixteConfig.Server.Port != float64(8080) {
		t.Error("Didn't read the specified port of 8080")
	}

	if dixteConfig.Database.Host != "localhost" {
		t.Error("Didn't read the specified host of the database")
	}

	if dixteConfig.Database.Port != float64(5432) {
		t.Error("Didn't read the specified port of the database")
	}

	if dixteConfig.Database.Name != "dixte_analytics" {
		t.Error("Didn't read the specified name of the database")
	}

	if dixteConfig.Database.Username != "luizpv9" {
		t.Error("Didn't read the specified username for the database")
	}

	if dixteConfig.Database.Password != "password" {
		t.Error("Didn't read the specified password for the database")
	}
}

func TestLoadFromFileWithMissingData(t *testing.T) {
	curDir := currentDirectory()
	dixteConfig, err := LoadFromFile(curDir + "/02_data.json")
	if err != nil {
		t.Error(err)
	}

	if dixteConfig.Server != nil {
		t.Error("The server config wasn't set to nil")
	}

	if dixteConfig.Database.Username != "" {
		t.Error("The username for database wasn't empty")
	}

	if dixteConfig.Database.Password != "" {
		t.Error("The password for database wasn't empty")
	}
}

func TestLoadFromFileWithBadFormat(t *testing.T) {
	curDir := currentDirectory()
	_, err := LoadFromFile(curDir + "/03_data.json")
	if err == nil {
		t.Error("An error wasn't detected while parsing the dixteConfig")
	}
}

func TestAssigningDefaultValues(t *testing.T) {
	curDir := currentDirectory()
	dixteConfig, err := LoadFromFile(curDir + "/02_data.json")
	if err != nil {
		t.Error(err)
	}
	dixteConfig.AssignDefaults()
	if dixteConfig.Server == nil {
		t.Error("Didn't assign default server")
	}

	if dixteConfig.Server.Port != float64(5002) {
		t.Error("Didn't copy the port from the default struct")
	}

	if dixteConfig.Database.Username != "postgres" {
		t.Error("Didn't copy the database username from the default struct")
	}

	if dixteConfig.Database.Password != "" {
		t.Error("Didn't copy the database password from the default struct")
	}
}
