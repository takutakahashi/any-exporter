package config

import (
	"os"

	"gopkg.in/yaml.v1"
)

type Config struct {
	Host       string          `yaml:"host"`
	Port       string          `yaml:"port"`
	Resolution string          `yaml:"resolution"`
	Metrics    []MetricsConfig `yaml:"metrics"`
}

type MetricsConfig struct {
	WorkDir    string   `yaml:"workdir"`
	ScriptPath string   `yaml:"scriptPath"`
	Resolution string   `yaml:"resolution"`
	Args       []string `yaml:"args"`
}

func Load(configPath string) (Config, error) {
	str, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = yaml.Unmarshal(str, &config)
	if err != nil {
		return Config{}, err
	}
	return FillDefaults(config), nil
}

func FillDefaults(config Config) Config {
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == "" {
		config.Port = "8080"
	}
	if config.Resolution == "" {
		config.Resolution = "1m"
	}
	for i, metric := range config.Metrics {
		if metric.Resolution == "" {
			config.Metrics[i].Resolution = config.Resolution
		}
	}
	return config
}
