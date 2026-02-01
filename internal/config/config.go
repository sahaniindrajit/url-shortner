package config

import (
	"log"
	"os"
)

type Config struct {
	Port    string
	Env     string
	BaseURL string
}

func Load() *Config {

	cfg := &Config{
		Port:    os.Getenv("PORT"),
		Env:     os.Getenv("ENV"),
		BaseURL: os.Getenv("BASE_URL"),
	}

	if cfg.Port == "" {
		log.Fatal("Port required")
	}
	if cfg.Env == "" {
		log.Fatal("Enviroment required")
	}
	if cfg.BaseURL == "" {
		log.Fatal("Base url requird")
	}

	return cfg
}
