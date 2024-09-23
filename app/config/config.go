package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	// Database
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBName     string `env:"DB_NAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBSSLMode  string `env:"DB_SSLMODE"`
}

var (
	once sync.Once
)

func LoadConfig() Config {
	cfg := Config{}

	once.Do(func() {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			log.Fatalf("Configuration cannot be read: %v\n", err)
		}
	})

	return cfg
}
