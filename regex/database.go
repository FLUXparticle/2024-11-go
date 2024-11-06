package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func setupDatabase() *sql.DB {
	// Alte Datenbankdatei löschen, falls sie existiert
	if _, err := os.Stat("./cocktails.db"); err == nil {
		err = os.Remove("./cocktails.db")
		if err != nil {
			log.Fatal("Fehler beim Löschen der alten Datenbank:", err)
		}
	}

	// SQLite-Datenbank erstellen
	db, err := sql.Open("sqlite3", "./cocktails.db")
	if err != nil {
		log.Fatal("Fehler beim Öffnen der Datenbank:", err)
	}

	// Tabellen erstellen
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS cocktails (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL
        );
        CREATE TABLE IF NOT EXISTS ingredients (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            cocktail_id INTEGER,
            amount TEXT,
            name TEXT,
            FOREIGN KEY (cocktail_id) REFERENCES cocktails(id)
        );
    `)
	if err != nil {
		log.Fatal("Fehler beim Erstellen der Tabellen:", err)
	}

	return db
}

func insertIntoDatabase(allCocktails []*Cocktail) {
	db := setupDatabase()
	defer db.Close()

	for _, cocktail := range allCocktails {
		result, err := db.Exec(`INSERT INTO cocktails (name) VALUES (?)`, cocktail.Name)
		if err != nil {
			log.Fatal(err)
		}
		cocktailID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		for _, instruction := range cocktail.Ingredients {
			// Zutat in die Datenbank einfügen
			_, err := db.Exec(`INSERT INTO ingredients (cocktail_id, amount, name) VALUES (?, ?, ?)`, cocktailID, instruction.AmountCL, instruction.Name)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
