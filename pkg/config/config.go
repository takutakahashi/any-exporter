package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

type Config struct {
	Host    string          `yaml:"host"`
	Port    string          `yaml:"port"`
	Metrics []MetricsConfig `yaml:"metrics"`
}

type MetricsConfig struct {
	WorkDir    string   `yaml:"workdir"`
	ScriptPath string   `yaml:"scriptPath"`
	Args       []string `yaml:"args"`
}

func Load(configPath string) (Config, error) {
	str, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = yaml.Unmarshal(str, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
