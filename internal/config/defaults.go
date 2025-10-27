package config

import "scripter/internal/script"

func NewDefault() *Config {
	return &Config{
		Predef: true,
		Scripts: []script.Script{

		},
	}
}
