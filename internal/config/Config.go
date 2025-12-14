package config

import (
	"os"
)

type DatabaseConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBPassword string
	DBUser     string
}

type AppConfig struct {
	JWTSecret string
	JWTTTL    int // in hours
}

func LoadDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBUser:     os.Getenv("DB_USER"),
	}
}
