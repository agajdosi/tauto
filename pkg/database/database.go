package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

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
	create table bots (id integer not null primary key, username text unique, password text, active integer);
	delete from bots;
	create table others (id integer not null primary key, username text unique, status text, active integer);
	delete from others;
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
