package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	// Database
	DBHost         string `env:"DB_HOST"`
	DBPort         string `env:"DB_PORT"`
	DBUser         string `env:"DB_USER"`
	DBName         string `env:"DB_NAME"`
	DBPassword     string `env:"DB_PASSWORD"`
	DBSSLMode      string `env:"DB_SSLMODE"`
	DBMAXConnPool  string `env:"DB_MAX_CONN_POOL" env-default:"20"`
	DBPoolLifetime string `env:"DB_POOL_LIFETIME" env-default:"300"` // seconds

	// Application
	Environment string `env:"ENV"`
}

func LoadConfig() *Config {
	cfg := Config{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration cannot be read: %v\n", err)
	}

	return &cfg
}
