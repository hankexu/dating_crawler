package model

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Saver   string   `json:"saver"`
	Workers []string `json:"workers"`
}

func LoadConfig(path string) (Config, error) {
	contents, err := ioutil.ReadFile(path)
	var config Config
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(contents, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
