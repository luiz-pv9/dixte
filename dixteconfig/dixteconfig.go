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
			SSLMode: "disable",
			Dbname: "dixte_analytics"
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
	Dbname                    string
	User                      string
	Password                  string
	Host                      string
	Port                      string
	SSLMode                   string
	Fallback_Application_Name string
	Connect_Timeout           string
	SSLCert                   string
	SSLKey                    string
	SSLRootCert               string
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
		return nil
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
