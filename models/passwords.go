package models

import (
	"database/sql"
)

const (
	insertHash = `INSERT INTO passwords ( user_id, hash ) VALUES ( $1, $2 );`
	getHash    = `SELECT hash FROM passwords WHERE user_id=$1;`
	updateHash = `UPDATE passwords SET (hash) = ($1) WHERE user_id=$2`
)

func InsertPasswordHash(db *sql.DB, username, hashPassword string) (err error) {
	var userID int64
	if err, userID = GetUserID(db, username); err != nil {
		return err
	}
	if _, err = db.Exec(insertHash, userID, hashPassword); err != nil {
		return err
	}
	return nil
}

func GetPasswordHash(db *sql.DB, username string) (err error, hash string) {
	_, ID := GetUserID(db, username)
	if err = db.QueryRow(getHash, ID).Scan(&hash); err != nil {
		return err, hash
	}
	return nil, hash
}

func UpdatePassword(db *sql.DB, username, newPassword string) (err error) {
	_, ID := GetUserID(db, username)
	if _, err = db.Exec(updateHash, newPassword, ID); err != nil {
		return err
	}
	return nil
}
