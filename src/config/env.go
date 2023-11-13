package config

import (
	"github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	ListenHost string `env:"LISTEN_HOST" envDefault:"0.0.0.0"`
	ListenPort string `env:"LISTEN_PORT" envDefault:"8080"`

	DatabaseName string `env:"DATABASE_NAME" envDefault:"./database.db"`

	LogLocation string `env:"LOG_LOCATION" envDefault:"/logs"`

	GinMode string `env:"GIN_MODE" envDefault:"release"`

	CorsAllowOrigins []string `env:"CORS_ALLOW_ORIGINS" envSeparator:"," envDefault:"http://localhost:5173"`
}

func Parse() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
