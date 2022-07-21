package models

import (
	"database/sql"
	"log"
)

const (
	insert           = `INSERT INTO users(username) VALUES ($1) RETURNING user_id;`
	selectByUsername = `SELECT * FROM users WHERE username=$1;`
	selectByID       = `SELECT * FROM users WHERE user_id=$1;`
	updateByUsername = `UPDATE users SET username=$1 WHERE username=$2;`
	updateByID       = `UPDATE users SET username=$1 WHERE user_id=$2;`
	deleteByUsername = `DELETE FROM users WHERE username=$1;`
	deleteByID       = `DELETE FROM users WHERE user_id=$1;`
)

func Insert(u Users) (userID int64, err error) {
	db := CreateConnection()
	err = db.QueryRow(insert, u.Username).Scan(&userID)
	if err != nil {
		log.Fatalln(err)
		return -1, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	return userID, nil
}

func SelectByUsername(username string) bool {
	db := CreateConnection()
	_, err := db.Exec(selectByUsername, username)
	if err != nil {
		log.Fatalln("Error selecting user by username:", err)
		return false
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	return true
}

// UpdateByUsername u is initial username && u2 is updated username
func UpdateByUsername(u, v string) bool {
	db := CreateConnection()
	_, err := db.Exec(updateByUsername, u, v)
	if err != nil {
		log.Fatalln("Error updating user by username:", err)
		return false
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	return true

}

func DeleteByUsername(u string) bool {
	db := CreateConnection()
	_, err := db.Exec(deleteByUsername, u)
	if err != nil {
		log.Fatalln("Error deleting user by username:", err)
		return false
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	return true
}

func SelectByID(id int64) bool {
	db := CreateConnection()
	_, err := db.Exec(selectByID, id)
	if err != nil {
		log.Fatalln("Error selecting user by username:", err)
		return false
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	return true
}

func UpdateByID(u Users) bool {
	db := CreateConnection()
	_, err := db.Exec(updateByID, u.Username)
	if err != nil {
		log.Fatalln("Error updating user by id:", err)
		return false
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	return true
}

func DeleteByID(id int64) bool {
	db := CreateConnection()
	_, err := db.Exec(deleteByID, id)
	if err != nil {
		log.Fatalln("Error deleting user by id:", err)
		return false
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	return true
}
