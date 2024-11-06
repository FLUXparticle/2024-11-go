package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// SetupDatabase öffnet die Datenbankverbindung und erstellt die Tabellen, falls sie nicht existieren.
func SetupDatabase() *sql.DB {
	// Öffnen der SQLite-Datenbank
	db, err := sql.Open("sqlite3", "./northwind.db")
	if err != nil {
		log.Fatal("Fehler beim Öffnen der Datenbank:", err)
	}

	return db
}
