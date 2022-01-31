package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type AppConfig struct {
	DBUser     string `env:"POSTGRES_USER"`
	DBPassword string `env:"POSTGRES_PASSWORD"`
	DBName     string `env:"POSTGRES_USER"`
}

var Env AppConfig

func init() {
	if err := env.Parse(&Env); err != nil {
		log.Printf("There was an error parsing environment variables: %s", err)
	}
}
