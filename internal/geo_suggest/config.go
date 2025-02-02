package geo_suggest

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type config struct {
	Port string `env:"SERVER_PORT"`
}

func LoadConfig() (config, error) {
	_, loaded := os.LookupEnv("SERVER_PORT")
	if !loaded {
		// when script runs in docker, we don't need to load environment variables from .env, because they're already loaded
		err := godotenv.Load("../../internal/geo_suggest/.env")
		if err != nil {
			return config{}, fmt.Errorf("Error initializing geo-suggest: %w", err)
		}
	}

	return env.ParseAs[config]()
}
