package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	APP_NAME    string
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSLMODE  string
}

func Load() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	cfg := &AppConfig{
		APP_NAME: getEnv("APP_NAME", "Expense Tracker"),
		PORT:     getEnv("PORT", "8080"),

		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_PORT:     getEnv("DB_PORT", "5432"),
		DB_USER:     getEnv("DB_USER", "postgres"),
		DB_PASSWORD: getEnv("DB_PASSWORD", "password"),
		DB_NAME:     getEnv("DB_NAME", "expense_tracker"),
		DB_SSLMODE:  getEnv("DB_SSLMODE", "disable"),
	}

	return cfg
}

func (c *AppConfig) GetDBConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", c.DB_HOST, c.DB_USER, c.DB_PASSWORD, c.DB_NAME, c.DB_PORT, c.DB_SSLMODE)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
