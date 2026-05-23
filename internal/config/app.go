package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	ProdEnv         = "production"
	DevEnv          = "development"
	DefaultAppPort  = "3000"
	DefaultLogLevel = "info"
	FallbackDsn     = "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
)

func init() {
	if env := os.Getenv("ENV"); env != ProdEnv {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file")
		}
	}
}

type AppConfig struct {
	Port     string
	LogLevel string
	Dsn      string
	Env      string
}

func buildDsn() string {
	env := os.Getenv("ENV")
	if env == DevEnv {
		if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
			return dsn
		}
	}
	if env == ProdEnv {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		name := os.Getenv("DB_NAME")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port) // TODO: check sslmode for prod
	}

	return FallbackDsn
}

func GetAppConfig() AppConfig {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = DefaultAppPort
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = DefaultLogLevel
	}

	return AppConfig{
		Port:     port,
		LogLevel: logLevel,
		Dsn:      buildDsn(),
		Env:      os.Getenv("ENV"),
	}
}
