package cmd

import (
	"2024-11-go/northwind/pkg/db"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Listet alle Kategorien auf",
	Long:  `Dieser Befehl zeigt alle Kategorien in der Northwind-Datenbank an.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Datenbank einrichten
		database := db.SetupDatabase()
		defer database.Close()

		// Kategorien abrufen
		categories, err := db.GetCategories(database)
		if err != nil {
			log.Fatal("Fehler beim Abrufen der Kategorien:", err)
		}

		// Kategorien anzeigen
		fmt.Println("Kategorien:")
		for id, name := range categories {
			fmt.Printf("%d: %s\n", id, name)
		}
	},
}

func init() {
	rootCmd.AddCommand(categoriesCmd)
}
