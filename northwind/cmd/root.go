package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "northwind",
	Short: "Northwind CLI ist ein Tool zur Abfrage der Northwind-Datenbank",
	Long:  `Northwind CLI ermöglicht es Benutzern, Kategorien, Produkte und Bestellmengen aus der Northwind-Datenbank über die Kommandozeile abzurufen.`,
}

// Execute führt den Root-Befehl aus.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
