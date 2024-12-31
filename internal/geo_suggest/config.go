package geo_suggest

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type config struct {
	Port string `env:"SERVER_PORT"`
}

func LoadConfig() (config, error) {
	err := godotenv.Load("../../internal/geo_suggest/.env")
	if err != nil {
		return config{}, fmt.Errorf("Error initializing geo-suggest: %w", err)
	}
	return env.ParseAs[config]()
}
