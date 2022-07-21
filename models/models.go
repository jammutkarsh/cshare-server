package models

import (
	"database/sql"
	"fmt"
	"github.com/JammUtkarsh/cshare-server/utils"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

type dbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type Users struct {
	UserID   int64
	Username string
}

type ClipStack struct {
	userID  int64
	clipID  int64
	message string
	secret  bool
}

func getConfig() *dbConfig {
	utils.LoadEnv(`./db.env`)
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	return &dbConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		Name:     dbName,
	}
}

func CreateConnection() *sql.DB {
	configs := getConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.Host, configs.Port, configs.Username, configs.Password, configs.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`\c cshare`)
	return db
}

func CloseConnection(db *sql.DB) {
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
}
