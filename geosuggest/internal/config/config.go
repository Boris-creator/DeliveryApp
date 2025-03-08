package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port         string `env:"SERVER_PORT"`
	DadataApiKey string `env:"DADATA_API_KEY"`
}

func LoadConfig() (Config, error) {
	_, loaded := os.LookupEnv("SERVER_PORT")
	if !loaded {
		// when script runs in docker, we don't need to load environment variables from .env, because they're already loaded
		err := godotenv.Load(".env")
		if err != nil {
			return Config{}, fmt.Errorf("initializing geo-suggest: %w", err)
		}
	}

	return env.ParseAs[Config]()
}
