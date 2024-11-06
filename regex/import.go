package main

import (
	"bufio"
	"log"
	"os"
)

func readAllCocktails() []*Cocktail {
	// Datei öffnen
	file, err := os.Open("cocktails.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// RegEx für Zutaten (mit und ohne Mengenangabe)
	//ingredientRegex := regexp.MustCompile(`((\d+)cl:)?(.+)`)

	var allCocktails []*Cocktail
	//var cocktail *Cocktail
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//line := scanner.Text()

	}

	return allCocktails
}
