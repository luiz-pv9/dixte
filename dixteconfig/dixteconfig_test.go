package dixteconfig

import (
	"os"
	"path/filepath"
	"strings"
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
	dixteConfig, err := LoadFromFile(filepath.Join(curDir, "test_data", "01_data.json"))
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

	if dixteConfig.App.Token_Size != float64(64) {
		t.Error("Didn't read the specified token size")
	}
}

func TestLoadFromFileWithMissingData(t *testing.T) {
	curDir := currentDirectory()
	dixteConfig, err := LoadFromFile(filepath.Join(curDir, "test_data", "02_data.json"))
	if err != nil {
		t.Error(err)
	}

	if dixteConfig.Server != nil {
		t.Error("The server config wasn't set to nil")
	}

	if dixteConfig.App != nil {
		t.Error("The app config wasn't set to nil")
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
	_, err := LoadFromFile(filepath.Join(curDir, "test_data", "03_data.json"))
	if err == nil {
		t.Error("An error wasn't detected while parsing the dixteConfig")
	}
}

func TestAssigningDefaultValues(t *testing.T) {
	curDir := currentDirectory()
	dixteConfig, err := LoadFromFile(filepath.Join(curDir, "test_data", "02_data.json"))
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

	if dixteConfig.App.Token_Size != float64(32) {
		t.Error("Didn't copy the token size from the default struct")
	}

	if dixteConfig.Database.Dbname != "dixte_analytics" {
		t.Error("Didn't copy the database name from the default struct")
	}

	if dixteConfig.Database.SSLMode != "disable" {
		t.Error("Didn't load default value for SSL mode connection")
	}
}

func TestDatabaseConfigToConnectionArguments(t *testing.T) {
	curDir := currentDirectory()
	dixteConfig, err := LoadFromFile(filepath.Join(curDir, "test_data", "01_data.json"))
	if err != nil {
		t.Error(err)
	}

	config := dixteConfig.Database.ToConnectionArguments()
	if strings.Index(config, "dbname=dixte_analytics") == -1 {
		t.Error("Didn't include the database name")
	}

	if strings.Index(config, "user=luizpv9") == -1 {
		t.Error("Didn't include the user")
	}

	if strings.Index(config, "password=password") == -1 {
		t.Error("Didn't include the password")
	}

	if strings.Index(config, "host=localhost") == -1 {
		t.Error("Didn't include the host")
	}

	if strings.Index(config, "port=5432") == -1 {
		t.Error("Didn't include the port")
	}

	if strings.Index(config, "sslmode=disable") == -1 {
		t.Error("Didn't include the sslmode")
	}

	if strings.Index(config, "fallback_application_name=dixte") == -1 {
		t.Error("Didn't include the fallback_application_name")
	}

	if strings.Index(config, "connect_timeout=20") == -1 {
		t.Error("Didn't include the connect_timeout")
	}

	if strings.Index(config, "sslcert=/home") == -1 {
		t.Error("Didn't include the sslcert")
	}

	if strings.Index(config, "sslkey=cat") == -1 {
		t.Error("Didn't include the sslkey")
	}

	if strings.Index(config, "sslrootcert=/root") == -1 {
		t.Error("Didn't include the sslrootcert")
	}
}

func TestDatabaseConfigToConnectionMissingArguments(t *testing.T) {
	curDir := currentDirectory()
	dixteConfig, err := LoadFromFile(filepath.Join(curDir, "test_data", "02_data.json"))
	if err != nil {
		t.Error(err)
	}

	config := dixteConfig.Database.ToConnectionArguments()

	if strings.Index(config, "dbname=dixte_analytics") == -1 {
		t.Error("Didn't include the database name")
	}

	if strings.Index(config, "user") != -1 {
		t.Error("Shouldn't include user")
	}

	if strings.Index(config, "password") != -1 {
		t.Error("Shouldn't include password")
	}

	if strings.Index(config, "host=localhost") == -1 {
		t.Error("Didn't include the host")
	}

	if strings.Index(config, "port=5432") == -1 {
		t.Error("Didn't include the port")
	}

	if strings.Index(config, "sslmode=disable") != -1 {
		t.Error("Shouldn't include the sslmode")
	}

	if strings.Index(config, "fallback_application_name=dixte") != -1 {
		t.Error("Shouldn't include the fallback_application_name")
	}

	if strings.Index(config, "connect_timeout=20") != -1 {
		t.Error("Shouldn't include the connect_timeout")
	}

	if strings.Index(config, "sslcert=/home") != -1 {
		t.Error("Shouldn't include the sslcert")
	}

	if strings.Index(config, "sslkey=cat") != -1 {
		t.Error("Shouldn't include the sslkey")
	}

	if strings.Index(config, "sslrootcert=/root") != -1 {
		t.Error("Shouldn't include the sslrootcert")
	}
}
