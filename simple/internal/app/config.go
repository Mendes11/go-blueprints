package app

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load() //nolint
	env := loadString("ENVIRONMENT", "development")
	godotenv.Load(fmt.Sprintf(".env.%s", env))
	godotenv.Load(fmt.Sprintf(".env.%s.local", env))
}

type Config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	Host        string `env:"HOST" envDefault:"127.0.0.1"`
	Port        int    `env:"PORT" envDefault:"5000"`
}

func LoadConfig() Config {
	c, err := env.ParseAs[Config]()
	if err != nil {
		panic(err)
	}
	return c
}

func loadString(name string, defaultValue string) string {
	if os.Getenv(name) != "" {
		return os.Getenv(name)
	}
	return defaultValue
}
