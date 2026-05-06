package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	defaultAppPort  = "8080"
	defaultLogLevel = "info"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
}

type AppConfig struct {
	Port     string
	LogLevel string
}

func GetAppConfig() AppConfig {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = defaultAppPort
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = defaultLogLevel
	}

	return AppConfig{
		Port:     port,
		LogLevel: logLevel,
	}
}
