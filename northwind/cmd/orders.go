package cmd

import (
	"2024-11-go/northwind/pkg/db"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var productID int

var ordersCmd = &cobra.Command{
	Use:   "orders",
	Short: "Zeigt die Gesamtbestellmenge eines Produkts an",
	Long:  `Dieser Befehl zeigt, wie oft ein ausgewähltes Produkt in der Northwind-Datenbank bestellt wurde.`,
	Run: func(cmd *cobra.Command, args []string) {
		if productID == 0 {
			fmt.Println("Bitte gib eine gültige Produkt-ID an. Verwende --product <ID>")
			return
		}

		// Datenbank einrichten
		database := db.SetupDatabase()
		defer database.Close()

		// Bestellmengen abrufen
		totalQuantity, err := db.GetTotalOrderQuantity(database, productID)
		if err != nil {
			log.Fatal("Fehler beim Abrufen der Bestellmengen:", err)
		}

		// Ergebnis anzeigen
		fmt.Printf("Das Produkt mit ID %d wurde insgesamt %d Mal bestellt.\n", productID, totalQuantity)
	},
}

func init() {
	rootCmd.AddCommand(ordersCmd)
	ordersCmd.Flags().IntVarP(&productID, "product", "p", 0, "ID des Produkts")
}
