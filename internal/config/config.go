package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type DbConfig struct {
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbServer   string `env:"DB_SERVER"`
}
type AppConfig struct {
	Host string `env:"HOST"`
	Port string `env:"PORT"`
	DbConfig
}

func LoadConfig() (AppConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return AppConfig{}, err
	}
	return env.ParseAs[AppConfig]()
}
