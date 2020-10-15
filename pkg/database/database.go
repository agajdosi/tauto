package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

//User holds all the data for the single user
type User struct {
	Username string
	Password string
	Platform string
}

//EnsureExists ensures that DB exists. If not, it setups a new database.
func EnsureExists() error {
	loc, err := Location()
	if err != nil {
		return err
	}

	_, err = os.Stat(loc)
	if err == nil {
		return nil
	} else if os.IsNotExist(err) {
		err = CreateDB(loc)
		if err != nil {
			return err
		}
		return nil
	}

	return err
}

//CreateDB creates a new database and sets all the needed tables.
func CreateDB(loc string) error {
	db, err := sql.Open("sqlite3", loc)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cmd := `
	create table bots (id integer not null primary key, username text unique, password text, platform text);
	delete from bots;
	`
	_, err = db.Exec(cmd)
	if err != nil {
		log.Printf("%q: %s\n", err, cmd)
		return err
	}

	return nil
}

//Location returns a location where the DB is located.
func Location() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(home, ".tst")
	err = os.MkdirAll(configDir, 0700)
	if err != nil {
		return "", err
	}

	location := filepath.Join(configDir, "tst.db")
	return location, nil
}

//AddBot adds a new bot into the database
func AddBot(username, password, platform string) error {
	location, err := Location()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	addCmd, err := tx.Prepare("insert into bots(username, password, platform) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer addCmd.Close()

	_, err = addCmd.Exec(username, password, platform)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}

//GetBot gets login information for the selected username.
func GetBot(username string) (string, error) {
	if username == "" {
		return "", fmt.Errorf("username was not not provided")
	}

	location, err := Location()
	if err != nil {
		return "", err
	}

	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return "", err
	}
	defer db.Close()

	stmt, err := db.Prepare("select password from bots where username = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var pass string
	err = stmt.QueryRow(username).Scan(&pass)
	if err != nil {
		log.Fatal(err)
	}

	return pass, nil
}

//GetAllBots gets login information of all bots from the database.
func GetAllBots() ([]User, error) {
	location, err := Location()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select username, password, platform from bots")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var username string
		var password string
		var platform string
		err = rows.Scan(&username, &password, &platform)
		if err != nil {
			return nil, err
		}
		u := User{username, password, platform}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		return users, err
	}

	return users, nil
}
