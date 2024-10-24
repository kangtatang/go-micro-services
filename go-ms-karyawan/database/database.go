package database

import (
	"database/sql"
	"log"

	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Connect to SQLite database
func Connect() {
	var err error
	DB, err = sql.Open("sqlite", "./employee.db")
	if err != nil {
		log.Fatal(err)
	}

	// Membuat tabel jika belum ada
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS employees (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        age INTEGER,
        position TEXT
    );`

	if _, err := DB.Exec(createTableQuery); err != nil {
		log.Fatal(err)
	}
}
