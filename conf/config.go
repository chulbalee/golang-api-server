package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const (
	SERVER_CONFIG_PATH = "conf/server.yaml"
)

type Config struct {
	Database struct {
		Type     string `yaml:"type"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func LoadServerConfig() *Config {
	config := Config{}

	configFile, err := ioutil.ReadFile(SERVER_CONFIG_PATH)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		panic(err)
	}

	return &config

}
