package db

import (
	"log"

	"github.com/joho/godotenv"
)

func GetKeysInEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}