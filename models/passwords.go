package models

import (
	"database/sql"
)

const (
	insertHash = `INSERT INTO passwords ( user_id, hash ) VALUES ( $1, $2 );`
	getHash    = `SELECT hash FROM passwords WHERE user_id=$1;`
	updateHash = `UPDATE passwords SET (hash) = ($1) WHERE user_id=$2`
)

func InsertPasswordHash(db *sql.DB, username, hashPassword string) (err error, val bool) {
	err, userID := GetUserID(db, username)
	if err != nil {
		return err, false
	}
	if _, err = db.Exec(insertHash, userID, hashPassword); err != nil {
		return err, false
	}
	return nil, true
}

func GetPasswordHash(db *sql.DB, username string) (err error, hash string) {
	_, ID := GetUserID(db, username)
	err = db.QueryRow(getHash, ID).Scan(&hash)
	if err != nil {
		return err, hash
	}
	return nil, hash
}

func UpdatePassword(db *sql.DB, username, newPassword string) (err error, val bool) {
	_, ID := GetUserID(db, username)
	_, err = db.Exec(updateHash, newPassword, ID)
	if err != nil {
		return err, false
	}
	return nil, true
}
