package models

// userQueries.go has methods which execute raw SQL statements in user table.
// Postgres version of SQL is being used here

import "database/sql"

const (
	insertUser       = `INSERT INTO users(username) VALUES ($1) RETURNING userID;`
	selectByUsername = `SELECT userID FROM users WHERE username=$1;`
)

// InsertUser inserts a new user into the database and returns a userID. userID is -1 if it already exists.
func InsertUser(db *sql.DB, username string) (userID int64, err error) {
	if err = db.QueryRow(insertUser, username).Scan(&userID); err != nil {
		return -1, err
	}
	return userID, nil
}

// SelectByUsername searches a user in DB by its username.
func SelectByUsername(db *sql.DB, username string) (userID int64, err error) {
	if err = db.QueryRow(selectByUsername, username).Scan(&userID); err != nil {
		return -1, err
	}
	return userID, nil
}

// GetUserID searches a user in DB by its userID. Uses username internally.
func GetUserID(db *sql.DB, username string) (userID int64, err error) {
	if userID, err = SelectByUsername(db, username); err != nil {
		return -1, err
	}
	return userID, nil
}
