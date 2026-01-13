package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Schema struct {
	Port                 string `env:"PORT"`
	PostgresUri          string `env:"POSTGRES_URI"`
	DisableAutoMigration bool   `env:"DISABLE_AUTO_MIGRATION" envDefault:"true"`
}

var cfg Schema

func GetConfig() *Schema {
	cfg.Port = os.Getenv("PORT")
	cfg.PostgresUri = os.Getenv("POSTGRES_URI")
	cfg.DisableAutoMigration = os.Getenv("DISABLE_AUTO_MIGRATION") == "true"
	if cfg.Port == "" {
		_ = godotenv.Load(filepath.Join(".env"))
		if err := env.Parse(&cfg); err != nil {
			log.Fatalf("Error on parsing configuration file, error: %v", err)
		}
	}
	return &cfg
}
