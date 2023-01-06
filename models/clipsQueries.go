package models

// clipsQueries.go has methods which execute raw SQL statements in user table.
// Postgres version of SQL is being used here

import (
	"database/sql"
)

const (
	insertClip = `INSERT INTO clip_stack ( userID, clipID, message, secret)
VALUES ($1, $2, $3, $4) RETURNING clipID`
	selectSingleClip = `SELECT clipID, message, secret FROM clip_stack WHERE clipID=$1 AND userID=$2`
	deleteSingleClip = `DELETE FROM clip_stack WHERE clipID=$1 AND userID=$2`
	deleteClips      = `DELETE FROM clip_stack WHERE userID=$1;`
	countClips       = `SELECT COUNT (userID) FROM clip_stack WHERE userID=$1 ;`
)

// ClipCount returns the number of clips for a user. Returns -1 if the user doesn't exist or has 0 clips.
func ClipCount(db *sql.DB, userID int64) (err error, count int64) {
	if err = db.QueryRow(countClips, userID).Scan(&count); err != nil {
		return err, -1
	}
	return nil, count
}

// InsertClip inserts a clip into the database and returns messageID.
func InsertClip(db *sql.DB, c Data) (err error, clipID int64) {
	if err, c.UserID = GetUserID(db, c.Username); err != nil {
		return err, -1
	}
	_, c.MessageID = ClipCount(db, c.UserID)
	if err = db.QueryRow(insertClip, c.UserID, c.MessageID+1, c.Message, c.Secret).Scan(&clipID); err != nil {
		return err, -1
	}
	return nil, clipID
}

// SelectClip returns clip Data for a given user.
func SelectClip(db *sql.DB, clipID, userID int64) (err error, c Data) {
	if err = db.QueryRow(selectSingleClip, clipID, userID).Scan(&c.MessageID, &c.Message, &c.Secret); err != nil {
		return err, c
	}
	c.UserID = userID
	return nil, c
}

func DeleteClip(db *sql.DB, clipID, userID int64) (err error) {
	if _, err = db.Exec(deleteSingleClip, clipID, userID); err != nil {
		return err
	}
	return nil
}

// DeleteClips deletes all the clips of a user.
func DeleteClips(db *sql.DB, userID int64) (err error) {
	if _, err = db.Exec(deleteClips, userID); err != nil {
		return err
	}
	return nil
}
