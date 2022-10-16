package utils

import (
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const (
	ColorReset  = string("\033[0m")
	ColorRed    = string("\033[31m")
	ColorGreen  = string("\033[32m")
	ColorYellow = string("\033[33m")
	ColorBlue   = string("\033[34m")
	ColorPurple = string("\033[35m")
	ColorCyan   = string("\033[36m")
	ColorWhite  = string("\033[37m")
)

// LoadEnv loads the .env file	and returns the error if any.
func LoadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Println(err)
		log.Fatalf("error loading .env file: %s", err.Error())
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}