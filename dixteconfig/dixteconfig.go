package dixteconfig

import (
	"encoding/json"
	"io/ioutil"
)

var (
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
