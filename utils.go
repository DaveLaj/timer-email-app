package main

import (
	"encoding/json"
	"os"
)

func LoadConfig() (Config, error) {
	var cfg Config
	var err error

	data, err := os.ReadFile("config.json")
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
