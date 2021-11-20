package data

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func check(e error) {
	fmt.Println(e)
}

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	check(err)

	return db.Ping()
}

func CreateTable() {
	createTableSQL := ``

	statement, err := db.Prepare(createTableSQL)
	check(err)

	statement.Exec()
	log.Println("Success on db connection")
}
