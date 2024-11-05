package gococktails

type Cocktail struct {
	CocktailID   uint           `json:"id"`
	Name         string         `json:"name"`
	Instructions []*Instruction `json:"instructions"`
}
type Instruction struct {
	AmountCL   int         `json:"amount"`
	Ingredient *Ingredient `json:"ingredient"`
}
type Ingredient struct {
	IngredientID uint   `json:"id"`
	Name         string `json:"name"`
}
