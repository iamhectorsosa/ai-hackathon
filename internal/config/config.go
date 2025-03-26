package config

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Environment string
}

func Load() *Config {
	cfg := &Config{
		Port:        "8080",
		Environment: "development",
	}

	envMap, err := godotenv.Read()
	if err != nil {
		log.Println("No .env file found, using default configuration")
		return cfg
	}

	if port := envMap["PORT"]; strings.TrimSpace(port) != "" {
		cfg.Port = port
	}

	if env := envMap["ENVIRONMENT"]; strings.TrimSpace(env) != "" {
		cfg.Environment = env
	}

	return cfg
}
