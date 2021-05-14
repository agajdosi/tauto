package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

//EnsureExists ensures that DB exists. If not, it setups a new database.
func EnsureExists() {
	loc, err := DBPath()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(loc)
	if err == nil {
		return
	} else if os.IsNotExist(err) {
		err = CreateDB(loc)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	return
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

//DBPath returns a location where the DB is located.
func DBPath() (string, error) {
	configDir := ConfigDirectory()
	location := filepath.Join(configDir, "tauto.db")

	return location, nil
}

//ConfigDirectory
func ConfigDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configDir := filepath.Join(home, ".tauto")
	err = os.MkdirAll(configDir, 0700)
	if err != nil {
		log.Fatal(err)
	}

	return configDir
}
