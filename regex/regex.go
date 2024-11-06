package main

type Cocktail struct {
	Name        string
	Ingredients []*Ingredient
}
type Ingredient struct {
	AmountCL int
	Name     string
}

func main() {
	allCocktails := readAllCocktails()
	insertIntoDatabase(allCocktails)
}
