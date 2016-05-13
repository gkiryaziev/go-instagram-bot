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
	Db        Db        `yaml:"db"`
	Instagram Instagram `yaml:"instagram"`
}

// Db struct
type Db struct {
	Driver   string `yaml:"driver"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

// Instagram struct
type Instagram struct {
	TimeoutMin int `yaml:"timeout_min"`
	TimeoutMax int `yaml:"timeout_max"`
	Users      []struct {
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
	} `yaml:"users"`
}
