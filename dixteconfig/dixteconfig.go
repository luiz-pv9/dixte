package dixteconfig

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"
)

var (
	// The values in the defaultConfig struct will be used in case they are
	// missing from the one the user loaded.
	defaultConfig = DixteConfig{
		Server: &ServerConfig{
			Port: "5002",
		},
		Database: &DatabaseConfig{
			SSLMode: "disable",
			Dbname:  "dixte_analytics",
		},
		App: &AppConfig{
			Token_Size: float64(32),
		},
	}
)

// The DixteConfig struct holds the data for initializing dixte analytics. It must
// have all configuration necessary to start the service, that means no other
// data must come from anywhere else.
type DixteConfig struct {
	Server   *ServerConfig
	Database *DatabaseConfig
	App      *AppConfig
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

func (d *DatabaseConfig) ToConnectionArguments() string {
	values := reflect.ValueOf(*d)
	fields := reflect.Indirect(values)
	config := ""
	for i := 0; i < values.NumField(); i++ {
		field := strings.ToLower(fields.Type().Field(i).Name)
		value := values.Field(i).String()
		if value != "" {
			config += field + "=" + values.Field(i).String()
			if i < values.NumField()-1 {
				config += " "
			}
		}
	}
	return config
}

type ServerConfig struct {
	Port string
}

type AppConfig struct {
	Token_Size float64
}

func (dc *DixteConfig) AssignDefaults() {
	dc.AssignServerDefaults()
	dc.AssignDatabaseDefaults()
	dc.AssignAppDefaults()
}

func (dc *DixteConfig) AssignServerDefaults() {
	if dc.Server == nil {
		// Maybe this should be copied instead of referenced.
		// At first look, there is no reason to change the values from the
		// DixteConfig struct in runtime.
		dc.Server = defaultConfig.Server
		return
	}
	if dc.Server.Port == "" {
		dc.Server.Port = defaultConfig.Server.Port
	}
}

func (dc *DixteConfig) AssignDatabaseDefaults() {
	if dc.Database == nil {
		dc.Database = defaultConfig.Database
		return
	}

	if dc.Database.SSLMode == "" {
		dc.Database.SSLMode = defaultConfig.Database.SSLMode
	}

	if dc.Database.Dbname == "" {
		dc.Database.Dbname = defaultConfig.Database.Dbname
	}
}

func (dc *DixteConfig) AssignAppDefaults() {
	if dc.App == nil {
		dc.App = defaultConfig.App
		return
	}

	if dc.App.Token_Size == float64(0) {
		dc.App.Token_Size = defaultConfig.App.Token_Size
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
