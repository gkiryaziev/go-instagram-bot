package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ConfigManager struct
type ConfigManager struct {
	file string
}

// NewConfig constructor
func NewConfig(file string) *ConfigManager {
	return &ConfigManager{file}
}

// Load config from file
func (cm *ConfigManager) Load() (*Config, error) {
	data, err := ioutil.ReadFile(cm.file)
	if err != nil {
		return nil, err
	}

	var config *Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Config struct
type Config struct {
	Debug     bool      `yaml:"debug"`
	RPC       RPC       `yaml:"rpc"`
	Db        Db        `yaml:"db"`
	Instagram Instagram `yaml:"instagram"`
}

// RPC struct
type RPC struct {
	AppHost string `yaml:"app_host"`
	AppPort int    `yaml:"app_port"`
}

// Db struct
type Db struct {
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_pass"`
	DbHost     string `yaml:"db_host"`
	DbPort     string `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
}

// Instagram struct
type Instagram struct {
	TimeoutMin int `yaml:"timeout_min"`
	TimeoutMax int `yaml:"timeout_max"`
	Users      []struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"users"`
}
