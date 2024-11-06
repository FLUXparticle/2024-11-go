package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readAllCocktails() []*Cocktail {
	// Datei öffnen
	file, err := os.Open("cocktails.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// RegEx für Zutaten (mit und ohne Mengenangabe)
	ingredientRegex := regexp.MustCompile(`((\d+)cl:)?(.+)`)

	var allCocktails []*Cocktail
	var cocktail *Cocktail
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if cocktail == nil {
			// Neuen Cocktail erstellen
			cocktail = &Cocktail{
				Name: line,
			}
		} else if ingredientRegex.MatchString(line) {
			// Zutaten-Zeilen verarbeiten
			matches := ingredientRegex.FindStringSubmatch(line)
			amount, err := strconv.Atoi(matches[2]) // Kann leer sein
			if err != nil {
				log.Println(err)
			}
			ingredient := matches[3]

			// Zutat in die Liste einfügen
			cocktail.Ingredients = append(cocktail.Ingredients, &Ingredient{
				AmountCL: amount,
				Name:     ingredient,
			})
		} else if line == "" { // Ende eines Cocktail-Blocks
			fmt.Println(cocktail.Name)
			for _, ingredient := range cocktail.Ingredients {
				fmt.Println(*ingredient)
			}
			cocktail = nil
		}
	}

	return allCocktails
}
