package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/JammUtkarsh/cshare-server/utils"
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
	UserID   int64
	Username string
}

type ClipStack struct {
	UserID  int64
	ClipID  int64
	Message string
	Secret  bool
}

func getConfig() *dbConfig {
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

// CreateConnection creates a connection to the database. It returns a pointer to the database.
// Don't forget to close the connection when you are done using it. Using CloseConnection() function.
func CreateConnection() *sql.DB {
	configs := getConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.Host, configs.Port, configs.Username, configs.Password, configs.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

// CloseConnection closes the connection to the database. It takes a pointer to the database as argument.
func CloseConnection(db *sql.DB) {
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
}
