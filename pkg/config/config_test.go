package config

import "testing"

func TestLoadConfig(t *testing.T) {
	config, err := Load("../../example/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if config.Host != "0.0.0.0" {
		t.Fatal(config)
	}
	if len(config.Metrics) != 1 {
		t.Fatal(config)
	}

	if config.Metrics[0].WorkDir != "example/test_metrics" {
		t.Fatal(config)
	}
}
