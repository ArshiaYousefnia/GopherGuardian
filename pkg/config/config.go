package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Targets []Target `json:"targets"`
}

type Alert struct {
	Email    string `json:"email"`
	Telegram string `json:"telegram"`
}

type Target struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Address  string `json:"address"`
	Interval int64  `json:"interval"`
	Alert    *Alert `json:"alert"`
}

func LoadConfig(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	err = validateConfig(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
