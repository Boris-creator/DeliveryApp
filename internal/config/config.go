package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type DbConfig struct {
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbServer   string `env:"DB_NAME"`
}
type AppConfig struct {
	Host           string `env:"HOST"`
	Port           string `env:"PORT"`
	GeoSuggestHost string `env:"GEO_SUGGEST_SERVER_HOST"`
	GeoSuggestPort string `env:"GEO_SUGGEST_SERVER_PORT"`
	DbConfig
}

var Config = AppConfig{}

func LoadConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("Error loading config: %w", err)
	}

	Config, err = env.ParseAs[AppConfig]()
	if err != nil {
		return fmt.Errorf("Error loading config: %w", err)
	}
	return nil
}
