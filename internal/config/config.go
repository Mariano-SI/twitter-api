package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	DbUrlMigration string `env:"DATABASE_URL"`
	DBPort         string `env:"DB_PORT"`
	DBUser         string `env:"DB_USER"`
	DBPassword     string `env:"DB_PASSWORD"`
	DBName         string `env:"DB_NAME"`
	DBHost         string `env:"DB_HOST"`
	JwtSecret      string `env:"JWT_SECRET"`
	Port           string `env:"PORT"`
}

func LoadConfig() (*config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, fmt.Errorf("failed to load .env file")
	}

	log.Println("config loaded")

	return &config{
		DbUrlMigration: os.Getenv("DATABASE_URL"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		DBHost:         os.Getenv("DB_HOST"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
		Port:           os.Getenv("PORT"),
	}, nil
}
