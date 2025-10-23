package config

import (
	"encoding/json"
	"scripter/internal/script"
)

type Config struct {
	Predef bool
	Scripts []script.Script
}

func NewLocalConfig(configData []byte) (*Config, error) {
	var cfg Config

	if err := json.Unmarshal(configData, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
