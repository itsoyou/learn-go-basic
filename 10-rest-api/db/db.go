package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
) // _ means we will not use this package directly, but go(database/sql built-in package) uses under the hood.

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // returns *sql.DB or error

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	/* configure connection pooling.
	 * How many open connections
	 * this database can have simultaneously.
	 * make sure it's not too big. */
	DB.SetMaxIdleConns(5)
	/* How many pools we want to keep open
	 * when no one is using it. */

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table.")
	}
}
