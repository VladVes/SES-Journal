package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const defaultAppPort = "8080"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

type appConf struct {
	Port string
}

func GetAppConfig() appConf {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = defaultAppPort
	}
	return appConf{
		Port: port,
	}
}
