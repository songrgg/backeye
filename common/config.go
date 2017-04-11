package common

import (
	"github.com/jinzhu/configor"
)

var Config = &ConfigBackeye{}

func init() {
	err := configor.Load(Config, "conf/backeye.yaml")
	if err != nil {
		panic("fail to load configuration")
	}
}

type ConfigBackeye struct {
	Bind      string    `json:"bind" yaml:"bind"`
	CertPem   string    `json:"cert_pem" yaml:"cert_pem"`
	KeyPem    string    `json:"key_pem" yaml:"key_pem"`
	Schedules Mongo     `json:"schedules"`
	Log       ConfigLog `json:"log"`
}

type Mongo struct {
	Address string `json:"address"`
	// DB      string `json:"db"`
}

// ConfigLog sets the logger level and destination
type ConfigLog struct {
	Level int `yaml:"level"`
}
