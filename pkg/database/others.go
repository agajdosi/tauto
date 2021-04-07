package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//Other holds all the data for the single other account (ally, neutral, enemy)
type Other struct {
	ID       int
	Username string
	Status   string
}

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

//GetOthers gets usernames of all active accounts of others (ally, neutral or enemy) from the database.
func GetOthers(username, status string) ([]Other, error) {
	location, err := Location()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var stmt *sql.Stmt
	//ideally add support for multiple usernames
	if username == "" {
		stmt, err = db.Prepare("select id, username, status from others where status = ? and active = 1")
	} else {
		stmt, err = db.Prepare("select id, username, status from others where username = ? and status = ? and active = 1")
	}

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows *sql.Rows
	if username != "" {
		rows, err = stmt.Query(username)
	} else {
		rows, err = stmt.Query()
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var others []Other
	for rows.Next() {
		var id int
		var username string
		var status string
		err = rows.Scan(&id, &username, &status)
		if err != nil {
			return nil, err
		}
		o := Other{id, username, status}
		others = append(others, o)
	}

	err = rows.Err()
	if err != nil {
		return others, err
	}

	return others, nil
}
