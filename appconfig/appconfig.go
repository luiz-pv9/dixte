package appconfig

import (
	"encoding/json"
	"io/ioutil"
)

// The AppConfig struct holds the data for initializing dixte analytics. It must
// have all configuration necessary to start the service, that means no other
// data must come from anywhere else.
type AppConfig struct {
	Port     float64
	Database struct {
		Host string
		Port float64
		Name string
	}
}

func LoadFromFile(filepath string) (*AppConfig, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	appConfig := AppConfig{}
	json.Unmarshal(content, &appConfig)
	return &appConfig, nil
}
