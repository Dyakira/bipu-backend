package middleware

import (
	"io/ioutil"
	"encoding/json"
)

var (
	Config *Configuration
)

const (
	ConfigFile = "src/main/conf/config.json"
)

type Db struct {
	User     string
	Pass     string
	Host     string
	Port     string
	Database string
}

// Credentials stores google client-ids.
type Credentials struct {
	ClientID     string `json:"clientid"`
	ClientSecret string `json:"secret"`
	RedirectURL string `json:"redirectUrl"`
}

type Session struct {
	SessionName string
	SessionSecret string
}


type Configuration struct {
	Db     Db
	Credentials Credentials
	Session Session
}

func GetConfig() (*Configuration, error) {
	var config Configuration
	if conf, err := ioutil.ReadFile(ConfigFile); err == nil {
		json.Unmarshal(conf, &config)
		return &config, nil
	} else {
		return &Configuration{}, err
	}
}