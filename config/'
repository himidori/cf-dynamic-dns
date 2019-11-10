package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	IPIdentURL string `json:"ipIdentUrl"`
	Sites      []Site `json:"sites"`
}

type Site struct {
	AuthEmail string `json:"authEmail"`
	AuthKey   string `json:"authKey"`
	ZoneID    string `json:"zoneId"`
	Domain    string `json:"domain"`
}

func GetConfig(path string) (*Config, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	if err := json.Unmarshal(body, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
