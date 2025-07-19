package config

import (
	"fmt"
	"os"
	"strconv"
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
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found, using system environment variables")
	// }

	config := &Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("PGHOST"),
			Port:     getEnvAsInt("PGPORT"),
			User:     os.Getenv("PGUSER"),
			Password: os.Getenv("PGPASSWORD"),
			DBName:   os.Getenv("PGDATABASE"),
			SSLMode:  os.Getenv("DB_SSL"),
			TimeZone: os.Getenv("DB_TIMEZONE"),
		},
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
			Host: os.Getenv("SERVER_HOST"),
		},
		API: APIConfig{
			APIKey: os.Getenv("API_KEY"),
		},
	}

	return config
}

func getEnvAsInt(key string) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return 0
}

func getEnvAsBool(key string) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return false
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
