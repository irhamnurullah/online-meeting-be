package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	API      APIConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	TimeZone string
	SSLMode  string
}

type ServerConfig struct {
	Port string
	Host string
}

type APIConfig struct {
	APIKey string
}

func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "123"),
			DBName:   getEnv("DB_NAME", "mydb"),
			SSLMode:  getEnv("DB_SSL", "disable"),
			TimeZone: getEnv("DB_TIMEZONE", "Asia/Jakarta"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Host: getEnv("SERVER_HOST", "localhost"),
		},
		API: APIConfig{
			APIKey: getEnv("API_KEY", ""),
		},
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Helper function untuk convert env var ke integer
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func (db *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		db.Host,
		db.User,
		db.Password,
		db.DBName,
		db.Port,
		db.SSLMode,
		db.TimeZone,
	)
}

// Helper function untuk convert env var ke boolean
// func getEnvAsBool(key string, defaultValue bool) bool {
// 	if value := os.Getenv(key); value != "" {
// 		if boolValue, err := strconv.ParseBool(value); err == nil {
// 			return boolValue
// 		}
// 	}
// 	return defaultValue
// }
