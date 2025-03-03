package config

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func Init() (*Cfg, error) {
	if err := load(); err != nil {
		return nil, fmt.Errorf("can't load config: %w", err)
	}

	cfg := &Cfg{}
	opts := env.Options{
		Prefix:                "APP_",
		UseFieldNameByDefault: true,
	}

	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, fmt.Errorf("can't parse config: %w", err)
	}

	return cfg, nil
}

func load() error { //nolint: unparam
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found: %v", err)
	}

	cfgEnv := os.Getenv("ENV_FILE")
	if len(cfgEnv) > 0 {
		err := godotenv.Load(cfgEnv)
		if err != nil {
			log.Fatalf("can't parse config: %v", err)
		}
	}
	return nil
}
