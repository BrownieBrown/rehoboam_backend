package env

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadLocalEnvironment() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env.local file")
	}
}

func LoadProductionEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env.prod file")
	}
}
