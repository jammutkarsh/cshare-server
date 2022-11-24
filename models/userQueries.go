package models

// userQueries.go has methods which execute raw SQL statements in user table.
// Postgres version of SQL is being used here

import "database/sql"

const (
	insertUser       = `INSERT INTO users(username) VALUES ($1) RETURNING user_id;`
	selectByUsername = `SELECT user_id FROM users WHERE username=$1;`
)

// InsertUser inserts a new user into the database and returns a userID. userID is -1 if it already exists.
func InsertUser(db *sql.DB, username string) (err error, userID int64) {
	if err = db.QueryRow(insertUser, username).Scan(&userID); err != nil {
		return err, -1
	}
	return nil, userID
}

// SelectByUsername searches a user in DB by its username.
func SelectByUsername(db *sql.DB, username string) (err error, userID int64) {
	if err = db.QueryRow(selectByUsername, username).Scan(&userID); err != nil {
		return err, -1
	}
	return nil, userID
}

// GetUserID searches a user in DB by its userID. Uses username internally.
func GetUserID(db *sql.DB, username string) (err error, userID int64) {
	if err, userID = SelectByUsername(db, username); err != nil {
		return err, -1
	}
	return nil, userID
}
