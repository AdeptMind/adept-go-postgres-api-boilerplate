package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	Port       int
	LogLevel   string
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbSslMode  string
	DbName     string
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Fatal("Error loading env vars")
	}
}

func GetConfig() *Config {
	return &Config{
		Port:       getEnvVarWithFallbackAsInt("PORT", 3000),
		LogLevel:   getEnvVarWithFallback("LOG_LEVEL", "info"),
		DbHost:     getEnvVarWithFallback("DB_HOST", "localhost"),
		DbPort:     getEnvVarWithFallbackAsInt("DB_PORT", 5432),
		DbUser:     getEnvVarWithFallback("DB_USER", "postgres"),
		DbPassword: getEnvVarWithFallback("DB_PASSWORD", ""),
		DbSslMode:  getEnvVarWithFallback("DB_SSL_MODE", "enable"),
		DbName:     getEnvVarWithFallback("DB_NAME", ""),
	}
}

func getEnvVarWithFallback(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvVarWithFallbackAsInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return intValue
}
