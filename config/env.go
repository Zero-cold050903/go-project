package config

import (
	"fmt"
	"os"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
}

func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func initConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "123456"),
		DBHost:     fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost:8080"), getEnv("DB_PORT", "8080")),
		DBName:     getEnv("DB_NAME", "go-project"),
	}
}
