package database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/glebarez/go-sqlite"
)

func createTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS entries (
        id INTEGER PRIMARY KEY,
        timestamp TEXT NOT NULL,
    	temperature REAL NOT NULL,
        gasLevel REAL NOT NULL
    );`

	return db.Exec(sql)
}

func SetupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	_, err = createTable(db)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	return db
}

func SetupDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./my.database")
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

func InsertEntry(db *sql.DB, entry DataEntry) (sql.Result, error) {
	sql := `insert into entries (timestamp, temperature, gasLevel) values (?, ?, ?)`
	result, err := db.Exec(sql, entry.TimeStamp, entry.Temperature, entry.GasLevel)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetEntries(db *sql.DB) ([]DataEntry, error) {
	sql := `select timestamp, temperature, gasLevel from entries`
	result, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var entries []DataEntry
	for result.Next() {
		var e DataEntry
		err := result.Scan(&e.TimeStamp, &e.Temperature, &e.GasLevel)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}
