package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("sqlite", "./project.db")
	if err != nil {
		log.Fatal(err)
	}

	// Buat tabel project jika belum ada
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS projects (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        description TEXT,
        employee_id INTEGER
    );`

	if _, err := DB.Exec(createTableQuery); err != nil {
		log.Fatal(err)
	}
}
