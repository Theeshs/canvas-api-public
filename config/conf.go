package config

import (
	"fmt"
	"log"
	"os"
)

// Config struct holds all app configuration
type Config struct {
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	AppPort       string
	OwnerEmail    string
	EmailPassword string
}

// AppConfig is the exported config instance
var AppConfig *Config

func LoadConfig() {
	AppConfig = &Config{
		DBUser:        getEnv("DB_USER", ""),
		DBPassword:    getEnv("DB_PASSWORD", ""),
		DBHost:        getEnv("DB_HOST", ""),
		DBPort:        getEnv("DB_PORT", ""),
		DBName:        getEnv("DB_NAME", ""),
		AppPort:       getEnv("APP_PORT", ""),
		OwnerEmail:    getEnv("EMAIL", ""),
		EmailPassword: getEnv("EMAIL_PASSWORD", ""),
	}
}

// Helper to read env variables or fallback to default
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// BuildDSN returns a PostgreSQL connection string
func (c *Config) BuildDSN() string {
	fmt.Printf("%s", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName,
	))
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName,
	)
}

// MustLoad ensures config is loaded or exits
func MustLoad() {
	LoadConfig()
	if AppConfig == nil {
		log.Fatal("failed to load config")
	}
}
