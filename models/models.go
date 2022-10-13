package models

import (
	"database/sql"
	"fmt"
	"github.com/JammUtkarsh/cshare-server/utils"
	"github.com/golang-jwt/jwt"
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

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
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
	// userID && clipID is the composite primary key of the table.
	// userID is the foreign key of the table.
	UserID   int64  `json:"userID"`
	Username string `json:"username" binding:"required"`
	// clipID is the is incremented by 1 every time a clip is added to the stack for each user.
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

// CreateConnection creates a connection to the database. It returns a pointer to the database.
// Don't forget to close the connection when you are done using it. Using CloseConnection() function.
func CreateConnection() *sql.DB {
	configs := getDBConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.Host, configs.Port, configs.Username, configs.Password, configs.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return db
}

// CloseConnection closes the connection to the database. It is already deferred. It takes a pointer to the database as argument.
func CloseConnection(db *sql.DB) {
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()
}
