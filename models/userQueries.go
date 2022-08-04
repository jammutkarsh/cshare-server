package models

import "database/sql"

const (
	insertUser       = `INSERT INTO users(username) VALUES ($1) RETURNING user_id;`
	selectByUsername = `SELECT user_id FROM users WHERE username=$1;`
	selectByUserID   = `SELECT username FROM users WHERE user_id=$1;`
	updateByUsername = `UPDATE users SET username=$1 WHERE username=$2;`
	updateByUserID   = `UPDATE users SET username=$1 WHERE user_id=$2;`
	deleteByUsername = `DELETE FROM users WHERE username=$1;`
	deleteByUserID   = `DELETE FROM users WHERE user_id=$1;`
)

// InsertUser inserts a new user into the database and returns the user_id of the new user.
//If the user already exists, it returns -1 and the error from DB.
func InsertUser(db *sql.DB, uname string) (err error, userID int64) {
	err = db.QueryRow(insertUser, uname).Scan(&userID)
	if err != nil {
		return err, -1
	}
	return nil, userID
}

//SelectByUsername checks if the user exists in database or not. If the user exists, it returns true.
//If the user does not exist, it returns false with the error from DB.
func SelectByUsername(db *sql.DB, uname string) (err error, userID int64) {
	err = db.QueryRow(selectByUsername, uname).Scan(&userID)
	if err != nil {
		return err, -1
	}
	return nil, userID
}

// UpdateByUsername changes the username of an already existing user. It returns true for a successful update of user.
//If the user does not exist, it returns false with the error from DB.
func UpdateByUsername(db *sql.DB, initialName, finalName string) (err error, val bool) {
	_, err = db.Exec(updateByUsername, finalName, initialName)
	if err != nil {
		return err, false
	}
	return nil, true

}

// DeleteByUsername deletes a user from the database. It returns true for a successful deletion of user.
//If the user does not exist, it returns false with the error from DB.
func DeleteByUsername(db *sql.DB, uname string) (err error, val bool) {
	_, err = db.Exec(deleteByUsername, uname)
	if err != nil {
		return err, false
	}
	return nil, true
}

// <- Miscellaneous functions for testing purposes. ->

// SelectByID checks if the user exists in database or not with user_id. If the user exists, it returns true.
//If the user does not exist, it returns false with the error from DB.
func SelectByID(db *sql.DB, userId int64) (err error, val bool) {
	_, err = db.Exec(selectByUserID, userId)
	if err != nil {
		return err, false
	}
	return nil, true
}

// UpdateByID changes the username of an already existing user with user_id. It returns true for a successful update of user.
//If the user does not exist, it returns false with the error from DB.
func UpdateByID(db *sql.DB, u Users) (err error, val bool) {
	_, err = db.Exec(updateByUserID, u.Username, u.UserID)
	if err != nil {
		return err, false
	}
	return nil, true
}

// DeleteByID deletes a user from the database with user_id. It returns true for a successful deletion of user.
//If the user does not exist, it returns false with the error from DB.
func DeleteByID(db *sql.DB, userId int64) (err error, val bool) {
	_, err = db.Exec(deleteByUserID, userId)
	if err != nil {
		return err, false
	}
	return nil, true
}
