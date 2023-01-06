package utils

// utils package has methods concerning common I/O operations to set up the environment.
import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file and returns the error if any.
func LoadEnv(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		log.Printf("error loading .env file: %s", err.Error())
		return err
	}
	return nil
}
