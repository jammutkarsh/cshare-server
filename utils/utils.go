package utils

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnv loads the .env file	and returns the error if any.
func LoadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("error loading db.env file: %s", err.Error())
	}
}
