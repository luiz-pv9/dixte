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

	if dixteConfig.Server.Port != "8080" {
		t.Error("Didn't read the specified port of 8080")
	}

	if dixteConfig.Database.Host != "localhost" {
		t.Error("Didn't read the specified host of the database")
	}

	if dixteConfig.Database.Port != "5432" {
		t.Error("Didn't read the specified port of the database")
	}

	if dixteConfig.Database.Dbname != "dixte_analytics" {
		t.Error("Didn't read the specified name of the database")
	}

	if dixteConfig.Database.User != "luizpv9" {
		t.Error("Didn't read the specified username for the database")
	}

	if dixteConfig.Database.Password != "password" {
		t.Error("Didn't read the specified password for the database")
	}

	if dixteConfig.Database.Connect_Timeout != "20" {
		t.Error("Didn't read the specified connect timeout")
	}

	if dixteConfig.Database.Fallback_Application_Name != "dixte" {
		t.Error("Didn't read the specified fallback application name")
	}

	if dixteConfig.Database.SSLCert != "/home" {
		t.Error("Didn't read the specified sslcert")
	}

	if dixteConfig.Database.SSLKey != "cat" {
		t.Error("Didn't read the specified sslkey")
	}

	if dixteConfig.Database.SSLRootCert != "/root" {
		t.Error("Didn't read the specified sslrootcert")
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

	if dixteConfig.Database.User != "" {
		t.Error("The username for database wasn't empty")
	}

	if dixteConfig.Database.Password != "" {
		t.Error("The password for database wasn't empty")
	}

	if dixteConfig.Database.SSLMode != "" {
		t.Error("The SSLMode for database wasn't empty")
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

	if dixteConfig.Server.Port != "5002" {
		t.Error("Didn't copy the port from the default struct")
	}

	if dixteConfig.Database.Dbname != "dixte_analytics" {
		t.Error("Didn't copy the database name from the default struct")
	}

	if dixteConfig.Database.SSLMode != "disable" {
		t.Error("Didn't load default value for SSL mode connection")
	}
}
