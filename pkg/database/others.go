package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//AddOther adds a new other account (ally, neutral or enemy) into the database
func AddOther(username, status string) (int, error) {
	location, err := Location()
	if err != nil {
		return -1, err
	}

	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return -1, err
	}

	addCmd, err := tx.Prepare("insert into others(username, status, active) values(?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer addCmd.Close()

	result, err := addCmd.Exec(username, status, 1)
	if err != nil {
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	return int(id), err
}
