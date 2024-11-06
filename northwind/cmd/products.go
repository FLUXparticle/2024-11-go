package cmd

import (
	"2024-11-go/northwind/pkg/db"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var categoryID int

var productsCmd = &cobra.Command{
	Use:   "products",
	Short: "Listet alle Produkte in einer Kategorie auf",
	Long:  `Dieser Befehl zeigt alle Produkte innerhalb einer angegebenen Kategorie in der Northwind-Datenbank an.`,
	Run: func(cmd *cobra.Command, args []string) {
		if categoryID == 0 {
			fmt.Println("Bitte gib eine gültige Kategorie-ID an. Verwende --category <ID>")
			return
		}

		// Datenbank einrichten
		database := db.SetupDatabase()
		defer database.Close()

		// Produkte abrufen
		products, err := db.GetProductsInCategory(database, categoryID)
		if err != nil {
			log.Fatal("Fehler beim Abrufen der Produkte:", err)
		}

		// Produkte anzeigen
		fmt.Printf("Produkte in Kategorie %d:\n", categoryID)
		for id, name := range products {
			fmt.Printf("%d: %s\n", id, name)
		}
	},
}

func init() {
	rootCmd.AddCommand(productsCmd)
	productsCmd.Flags().IntVarP(&categoryID, "category", "c", 0, "ID der Kategorie")
}
