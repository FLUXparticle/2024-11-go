package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func ExampleXML() {
	cocktail := Cocktail{
		Name: "Cranberry Cooler",
		Instructions: []Instruction{
			{Cl: 10, Text: "Preiselbeernektar"},
			{Cl: 5, Text: "Apfelsaft"},
			{Text: "Zitronensaft"},
			{Text: "Mineralwasser mit Kohlensäure"},
		},
	}

	// XML-Darstellung in indented format erstellen und ausgeben
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(cocktail); err != nil {
		fmt.Printf("Fehler beim XML-Encoding: %v\n", err)
	}

	// Output:
	// <cocktail name="Cranberry Cooler">
	//   <inst cl="10">Preiselbeernektar</inst>
	//   <inst cl="5">Apfelsaft</inst>
	//   <inst>Zitronensaft</inst>
	//   <inst>Mineralwasser mit Kohlensäure</inst>
	// </cocktail>
}
