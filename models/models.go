package models

// models package deals with all the database related queries and functions.
// models.go consists of
// 1. Data structure concerned with db connection, user and clip data.
// 2. Functions concerning database connection.

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type dbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type Users struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"  binding:"required"`
	PCount   int    `json:"pCount"  binding:"required"`
	SPCount  int    `json:"spCount"  binding:"required"`
}

type Data struct {
	UserID    int64  `json:"userID"`
	Username  string `json:"username"`
	MessageID int64  `json:"clipID"`
	Message   string `json:"message" binding:"required"`
	Secret    *bool  `json:"secret" binding:"required"`
}

// getDBConfig reads .env file for database authentication.
func getDBConfig() *dbConfig {
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &dbConfig{
		Port:     dbPort,
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_DATABASE"),
	}
}

// CreateConnection creates a connection to the database. Add CloseConnection in the next line.
func CreateConnection() (db *sql.DB, err error) {
	configs := getDBConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.Host, configs.Port, configs.Username, configs.Password, configs.Name)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}

// CloseConnection closes the connection to the database. Use defer CloseConnection below CreateConnection.
func CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}
