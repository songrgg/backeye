package std

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/coreos/etcd/clientv3"
	"github.com/jinzhu/configor"
	yaml "gopkg.in/yaml.v2"
)

var Config = &ConfigBackeye{}

func init() {
	path := "conf/backeye.yaml"
	if os.Getenv("CONFIG_ETCDS") != "" {
		loadFromEtcd(Config, path)
		return
	}

	err := configor.Load(Config, path)
	if err != nil {
		panic("fail to load configuration")
	}
}

type ConfigBackeye struct {
	Bind      string      `json:"bind" yaml:"bind"`
	CertPem   string      `json:"cert_pem" yaml:"cert_pem"`
	KeyPem    string      `json:"key_pem" yaml:"key_pem"`
	Schedules Mongo       `json:"schedules" yaml:"schedules"`
	MySQL     ConfigMySQL `yaml:"mysql"`
	Log       ConfigLog   `yaml:"log"`
}

type Mongo struct {
	Address string `json:"address"`
	// DB      string `json:"db"`
}

type ConfigMySQL struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxConn  int    `yaml:"max_conn"`
	LogType  string `yaml:"log_type"`
}

// ConfigLog sets the logger level and destination
type ConfigLog struct {
	Level int `yaml:"level"`
}

// ConfEnv fetches the current runtime environment
func ConfEnv() string {
	return configor.ENV()
}

func loadFromEtcd(dest interface{}, path string) {
	etcdList := os.Getenv("CONFIG_ETCDS")
	if etcdList == "" {
		LogPanic(LogFields{
			TagCategory: "config",
			"error":     "etcd list is empty",
		}, "fail to load configuration")
	}

	path = getConfigurationWithENV(path, ConfEnv())
	LogInfo(LogFields{
		TagCategory: "config",
	}, "start to load configuration "+path)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(etcdList, ","),
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		LogPanic(LogFields{
			TagCategory: "config",
			"error":     err,
		}, "fail to load configuration")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	rsp, err := cli.Get(ctx, path)
	cancel()
	if err != nil {
		LogPanic(LogFields{
			TagCategory: "config",
			"error":     err,
		}, "fail to load configuration")
	}

	err = load(dest, path, rsp.Kvs[0].Value)
	if err != nil {
		LogPanic(LogFields{
			TagCategory: "config",
			"error":     err,
		}, "fail to load configuration")
	}
}

func load(config interface{}, file string, data []byte) (err error) {
	switch {
	case strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml"):
		return yaml.Unmarshal(data, config)
	case strings.HasSuffix(file, ".toml"):
		return toml.Unmarshal(data, config)
	case strings.HasSuffix(file, ".json"):
		return json.Unmarshal(data, config)
	default:
		if toml.Unmarshal(data, config) != nil {
			if json.Unmarshal(data, config) != nil {
				if yaml.Unmarshal(data, config) != nil {
					return errors.New("failed to decode config")
				}
			}
		}
		return nil
	}
}

func getConfigurationWithENV(file, env string) string {
	var envFile string
	var extname = path.Ext(file)

	if extname == "" {
		envFile = fmt.Sprintf("%v.%v", file, env)
	} else {
		envFile = fmt.Sprintf("%v.%v%v", strings.TrimSuffix(file, extname), env, extname)
	}
	return filepath.Base(envFile)
}
