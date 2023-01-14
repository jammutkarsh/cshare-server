package models

// passwords.go has methods which execute raw SQL statements in passwords table.
// Postgres version of SQL is being used here

import (
	"database/sql"
)

const (
	insertHash = `INSERT INTO passwords ( userid, hash ) VALUES ( $1, $2 );`
	getHash    = `SELECT hash FROM passwords WHERE userid=$1;`
)

// InsertPasswordHash inserts the password hash in the database.
func InsertPasswordHash(db *sql.DB, username, hashPassword string) (err error) {
	var userID int64
	if userID, err = GetUserID(db, username); err != nil {
		return err
	}
	if _, err = db.Exec(insertHash, userID, hashPassword); err != nil {
		return err
	}
	return nil
}

// GetPasswordHash fetches the hash from the database.
func GetPasswordHash(db *sql.DB, username string) (hash string, err error) {
	ID, err := GetUserID(db, username)
	if err != nil {
		return "", err
	}
	if err = db.QueryRow(getHash, ID).Scan(&hash); err != nil {
		return hash, err
	}
	return hash, nil
}
