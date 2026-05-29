package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUrlMigration    string        `env:"DATABASE_URL"`
	DBPort            string        `env:"DB_PORT"`
	DBUser            string        `env:"DB_USER"`
	DBPassword        string        `env:"DB_PASSWORD"`
	DBName            string        `env:"DB_NAME"`
	DBHost            string        `env:"DB_HOST"`
	JwtSecret         string        `env:"JWT_SECRET"`
	RefreshTokenTTL   time.Duration `env:"REFRESH_TOKEN_TTL"`
	Port              string        `env:"PORT"`
	R2AccountID       string        `env:"R2_ACCOUNT_ID"`
	R2AccessKeyID     string        `env:"R2_ACCESS_KEY_ID"`
	R2SecretAccessKey string        `env:"R2_SECRET_ACCESS_KEY"`
	R2Bucket          string        `env:"R2_BUCKET"`
	R2PublicURL       string        `env:"R2_PUBLIC_URL"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, fmt.Errorf("failed to load .env file")
	}

	ttl, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_TTL"))
	if err != nil {
		return nil, fmt.Errorf("invalid REFRESH_TOKEN_TTL: %w", err)
	}

	log.Println("config loaded")

	return &Config{
		DbUrlMigration:    os.Getenv("DATABASE_URL"),
		DBPort:            os.Getenv("DB_PORT"),
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            os.Getenv("DB_NAME"),
		DBHost:            os.Getenv("DB_HOST"),
		JwtSecret:         os.Getenv("JWT_SECRET"),
		Port:              os.Getenv("PORT"),
		RefreshTokenTTL:   ttl,
		R2AccountID:       os.Getenv("R2_ACCOUNT_ID"),
		R2AccessKeyID:     os.Getenv("R2_ACCESS_KEY_ID"),
		R2SecretAccessKey: os.Getenv("R2_SECRET_ACCESS_KEY"),
		R2Bucket:          os.Getenv("R2_BUCKET"),
		R2PublicURL:       os.Getenv("R2_PUBLIC_URL"),
	}, nil
}
