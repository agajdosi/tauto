package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//AddBot adds a new bot into the database
func AddBot(username, password string) (int, error) {
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

	addCmd, err := tx.Prepare("insert into bots(username, password, active) values(?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer addCmd.Close()

	result, err := addCmd.Exec(username, password, 1)
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

//GetBots gets login information of all active bots from the database.
func GetBots(username string, active bool) ([]User, error) {
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
	if username != "" && active == true {
		//ideally add support for multiple usernames
		stmt, err = db.Prepare("select id, username, password from bots where username = ? and active = 1")
	} else if username != "" && active == false {
		stmt, err = db.Prepare("select id, username, password from bots where username = ?")
	} else if username == "" && active == true {
		stmt, err = db.Prepare("select id, username, password from bots where active = 1")
	} else {
		stmt, err = db.Prepare("select id, username, password from bots")
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

	var users []User
	for rows.Next() {
		var id int
		var username string
		var password string
		err = rows.Scan(&id, &username, &password)
		if err != nil {
			return nil, err
		}
		u := User{id, username, password}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		return users, err
	}

	return users, nil
}
