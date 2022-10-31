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
	// userID is the primary key of the table. It is autoincrement.
	UserID   int64  `json:"userID"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password"  binding:"required"`
	PCount   int    `json:"pCount"  binding:"required"`
	SPCount  int    `json:"spCount"  binding:"required"`
}

type Data struct {
	UserID    int64  `json:"userID"`
	Username  string `json:"username" binding:"required"`
	MessageID int64  `json:"clipID"`
	Message   string `json:"message" binding:"required"`
	Secret    bool   `json:"secret" binding:"required"`
}

func getDBConfig() *dbConfig {
	utils.LoadEnv(`./.env`)
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

// CreateConnection creates a connection to the database. Add CloseConnection() in the next line.
func CreateConnection() *sql.DB {
	configs := getDBConfig()
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		configs.Username, configs.Password, configs.Host, configs.Port, configs.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return db
}

// CloseConnection closes the connection to the database. It is already deferred. So that close connection pairs with create connection.
// Leaving no room for closing the connection, later.
func CloseConnection(db *sql.DB) {
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()
}
