package models

import "database/sql"

const (
	insertUser       = `INSERT INTO users(username) VALUES ($1) RETURNING user_id;`
	selectByUsername = `SELECT user_id FROM users WHERE username=$1;`
)

// InsertUser inserts a new user into the database and returns a userID. userID is -1 if it already exists.
func InsertUser(db *sql.DB, uname string) (err error, userID int64) {
	if err = db.QueryRow(insertUser, uname).Scan(&userID); err != nil {
		return err, -1
	}
	return nil, userID
}

// SelectByUsername checks for the existence of a user. Returns -1 if the user doesn't exist
func SelectByUsername(db *sql.DB, uname string) (err error, userID int64) {
	if err = db.QueryRow(selectByUsername, uname).Scan(&userID); err != nil {
		return err, -1
	}
	return nil, userID
}

// GetUserID checks for the existence of a user. Returns -1 if the user doesn't exist
func GetUserID(db *sql.DB, username string) (err error, userID int64) {
	if err, userID = SelectByUsername(db, username); err != nil {
		return err, -1
	}
	return nil, userID
}
