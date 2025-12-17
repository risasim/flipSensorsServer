package main

import (
	"database/sql"
	"fmt"
)

func createTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS entries (
        id INTEGER PRIMARY KEY,
        timestampe TEXT NOT NULL
    	temperature REAL NOT NULL,
        gasLevel REAL NOT NULL
    );`

	return db.Exec(sql)
}

func SetupDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./my.db")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = createTable(db)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var sqliteVersion string
	err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(sqliteVersion)

	return db, nil
}

func insertEntry(db *sql.DB, entry DataEntry) (sql.Result, error) {
	sql := `insert into entries (timestamp, temperature, gasLevel) values (?, ?, ?)`
	result, err := db.Exec(sql, entry.temperature, entry.gasLevel)
	if err != nil {
		return nil, err
	}
	return result, nil
}
