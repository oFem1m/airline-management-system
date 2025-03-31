package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config содержит параметры конфигурации приложения.
type Config struct {
	ServerAddress string
	DatabaseURL   string
}

// LoadConfig загружает конфигурацию из .env или переменных окружения.
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден, используем переменные окружения")
	}

	serverAddress := getEnv("SERVER_ADDRESS", ":8080")

	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "Airline")

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	cfg := &Config{
		ServerAddress: serverAddress,
		DatabaseURL:   databaseURL,
	}

	return cfg, nil
}

// getEnv получает переменную окружения или возвращает значение по умолчанию.
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
