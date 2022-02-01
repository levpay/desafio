package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type AppConfig struct {
	DBUser     string `env:"POSTGRES_USER"`
	DBPassword string `env:"POSTGRES_PASSWORD"`
	DBName     string `env:"POSTGRES_USER"`
	ServerPort string `env:"SERVER_PORT" envDefault:"3000"`
	APIKey     string `env:"API_KEY" envDefault:"4745594135527281"`
}

var Env AppConfig

func init() {
	if err := env.Parse(&Env); err != nil {
		log.Printf("There was an error parsing environment variables: %s", err)
	}
}
