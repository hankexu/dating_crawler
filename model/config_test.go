package model

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("config_test.json")
	if err != nil {
		panic(err)
	}
	if config.Saver != "1.1.1.1:2000" {
		t.Errorf("Error config: %v", config)
	}

}
