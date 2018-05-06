package middleware

import (
	"encoding/json"
	"io/ioutil"
)

var (
	Config *Configuration
)

const (
	ConfigFile = "src/main/conf/config.json"
	// ConfigFile = "conf/config.json"
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
	RedirectURL  string `json:"redirectUrl"`
}

type Session struct {
	SessionName   string
	SessionSecret string
}

type Others struct {
	Md5Salt string
}

type Configuration struct {
	Db          Db
	Credentials Credentials
	Session     Session
	Others      Others
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
