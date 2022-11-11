package models

import "database/sql"

const (
	insertUser       = `INSERT INTO users(username) VALUES ($1) RETURNING user_id;`
	selectByUsername = `SELECT user_id FROM users WHERE username=$1;`
)

// InsertUser inserts a new user into the database and returns the user_id of the new user.
// If the user already exists, it returns -1 and the error from DB.
func InsertUser(db *sql.DB, uname string) (err error, userID int64) {
	err = db.QueryRow(insertUser, uname).Scan(&userID)
	if err != nil {
		return err, -1
	}
	return nil, userID
}

// SelectByUsername checks if the user exists in database or not. If the user exists, it returns true.
// If the user does not exist, it returns false with the error from DB.
func SelectByUsername(db *sql.DB, uname string) (err error, userID int64) {
	err = db.QueryRow(selectByUsername, uname).Scan(&userID)
	if err != nil {
		return err, -1
	}
	return nil, userID
}
