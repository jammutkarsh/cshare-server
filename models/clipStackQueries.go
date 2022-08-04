package models

import (
	"database/sql"
	"log"
)

const (
	insertClip = `INSERT INTO clip_stack ( user_id, clip_id, message, secret)
VALUES ($1, $2, $3, $4) RETURNING clip_id;`
	selectClip = `SELECT clip_id, message, secret FROM clip_stack WHERE clip_id=$1 AND user_id=$2;`
	deleteClip = `DELETE FROM clip_stack WHERE clip_id=$1 AND user_id=$2;`
	countClips = `SELECT COUNT (user_id) FROM clip_stack WHERE user_id=$1 ;`
)

// GetUserID checks if the user exists, if it does, returns user_id of the user. If not, creates a new user and returns user_id of the new user.
func GetUserID(db *sql.DB, username string) (err error, userID int64) {
	err, userID = SelectByUsername(db, username)
	if err != nil {
		return err, -1
	}
	return nil, userID
}

// ClipCount returns the number of clips in the DB for a user.
func ClipCount(db *sql.DB, userID int64) (err error, count int64) {
	err = db.QueryRow(countClips, userID).Scan(&count)
	if err != nil {
		return err, -1
	}
	return nil, count
}

// InsertClip inserts a new clip into the database and returns the clip_id of the new clip.
func InsertClip(db *sql.DB, c ClipStack, username string) (err error, clipID int64) {
	err, c.UserID = GetUserID(db, username)
	if err != nil {
		return err, -1
	}
	err, c.ClipID = ClipCount(db, c.UserID)
	err = db.QueryRow(insertClip, c.UserID, c.ClipID+1, c.Message, c.Secret).Scan(&clipID)
	if err != nil {
		log.Fatalln("inside insertClip", err)
		return err, -1
	}
	return nil, clipID
}

// SelectClip returns the clip with the given clip_id and user_id.
func SelectClip(db *sql.DB, clipID, userID int64) (err error, c ClipStack) {
	err = db.QueryRow(selectClip, clipID, userID).Scan(&c.ClipID, &c.Message, &c.Secret)
	if err != nil {
		// TODO: handle error when no clip or user is found.
		return err, c
	}
	c.UserID = userID
	return nil, c
}

// DeleteClip deletes the clip with the given clip_id and user_id.
func DeleteClip(db *sql.DB, clipID, userID int64) (err error) {
	_, err = db.Exec(deleteClip, clipID, userID)
	// TODO: handle error when there is no clip to delete.
	if err != nil {
		return err
	}
	return nil
}
