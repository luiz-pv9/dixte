package dixteconfig

import (
	"encoding/json"
	"io/ioutil"
)

var (
	// The values in the defaultConfig struct will be used in case they are
	// missing from the one the user loaded.
	defaultConfig = DixteConfig{
		Server: &ServerConfig{
			Port: float64(5002),
		},
		Database: &DatabaseConfig{
			Host:     "localhost",
			Port:     float64(5432),
			Name:     "dixte_analytics",
			Username: "postgres",
			Password: "",
		},
	}
)

// The DixteConfig struct holds the data for initializing dixte analytics. It must
// have all configuration necessary to start the service, that means no other
// data must come from anywhere else.
type DixteConfig struct {
	Server   *ServerConfig
	Database *DatabaseConfig
}

// DatabaseConfig struct deals with configuration for interacting with
// PostgreSQL.
type DatabaseConfig struct {
	Host     string
	Port     float64
	Name     string
	Username string
	Password string
}

type ServerConfig struct {
	Port float64
}

func (dc *DixteConfig) AssignDefaults() {
	dc.AssignServerDefaults()
	dc.AssignDatabaseDefaults()
}

func (dc *DixteConfig) AssignServerDefaults() {
	if dc.Server == nil {
		// Maybe this should be copied instead of referenced.
		// At first look, there is no reason to change the values from the
		// DixteConfig struct in runtime.
		dc.Server = defaultConfig.Server
		return
	}
	if dc.Server.Port == float64(0) {
		dc.Server.Port = defaultConfig.Server.Port
	}
}

func (dc *DixteConfig) AssignDatabaseDefaults() {
	if dc.Database == nil {
		dc.Database = defaultConfig.Database
		return
	}

	if dc.Database.Host == "" {
		dc.Database.Host = defaultConfig.Database.Host
	}

	if dc.Database.Name == "" {
		dc.Database.Name = defaultConfig.Database.Name
	}

	if dc.Database.Username == "" {
		dc.Database.Username = defaultConfig.Database.Username
	}

	if dc.Database.Password == "" {
		dc.Database.Password = defaultConfig.Database.Password
	}

	if dc.Database.Port == float64(0) {
		dc.Database.Port = defaultConfig.Database.Port
	}
}

func LoadFromFile(filepath string) (*DixteConfig, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	dixteConfig := DixteConfig{}
	if err := json.Unmarshal(content, &dixteConfig); err != nil {
		return nil, err
	}
	return &dixteConfig, nil
}
