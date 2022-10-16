package models

import "database/sql"

func GetPasswordHash(db *sql.DB, username string) string {

	// db operations to get the hashPassword from the database
	return ""
}

func InsertPasswordHash(db *sql.DB, username, hashPassword string) {
	// db operations to insert the hashPassword into the database
}

func UpdatePassword(db *sql.DB, username, newPassword string) {
	// db operations to update the password
}
