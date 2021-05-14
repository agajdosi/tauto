package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//Bot holds all the data for the single bot
type Bot struct {
	ID       int
	Username string
	Password string
}

//AddBot adds a new bot into the database
func AddBot(username, password string) (int, error) {
	location, err := DBPath()
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
func GetBots(username string, onlyActive bool) ([]Bot, error) {
	location, err := DBPath()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var stmt *sql.Stmt
	if username != "" && onlyActive == true {
		//ideally add support for multiple usernames
		stmt, err = db.Prepare("select id, username, password from bots where username = ? and active = 1")
	} else if username != "" && onlyActive == false {
		stmt, err = db.Prepare("select id, username, password from bots where username = ?")
	} else if username == "" && onlyActive == true {
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

	var bots []Bot
	for rows.Next() {
		var id int
		var username string
		var password string
		err = rows.Scan(&id, &username, &password)
		if err != nil {
			return nil, err
		}
		b := Bot{id, username, password}
		bots = append(bots, b)
	}

	err = rows.Err()
	if err != nil {
		return bots, err
	}

	return bots, nil
}

//ListBots lists all bots
func ListBots() error {
	bots, err := GetBots("", true)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if len(bots) == 0 {
		fmt.Println("There aren't any bots in the database.")
		return nil
	}

	fmt.Println("Available bot accounts are:")
	for _, bot := range bots {
		fmt.Println(bot.ID, bot.Username, bot.Password)
	}

	return nil
}

//DeleteBot deletes a bot from the database.
func DeleteBot(botName string) error {
	location, err := DBPath()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM bots WHERE username = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(botName)
	deleted, _ := result.RowsAffected()
	fmt.Printf("Bots deleted: %v\n", deleted)

	return err
}
