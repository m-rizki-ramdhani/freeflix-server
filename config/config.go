package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_NAME  string
	GRPC_PORT string
	REST_PORT string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed get env file, using default value")
	}

	return Config{
		APP_NAME:  getEnv("APP_NAME", "APP_NAME"),
		GRPC_PORT: getEnv("GRPC_PORT", ":5001"),
		REST_PORT: getEnv("REST_PORT", ":5000"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}
