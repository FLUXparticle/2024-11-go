package gococktails

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getCocktails() []*Cocktail {
	resp, err := http.Get("https://cocktails.fluxparticle.com/api/cocktails")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data []*Cocktail // []map[string]any
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data
}

func TestSequential(t *testing.T) {
	cocktails := getCocktails()

	sum := 0

	for _, cocktail := range cocktails {
		sum += len(cocktail.Instructions)
	}

	assert.Equal(t, 38, sum)
}

func TestParallel(t *testing.T) {
	cocktails := getCocktails()

	sum := 0

	for _, cocktail := range cocktails {
		sum += len(cocktail.Instructions)
	}

	assert.Equal(t, 38, sum)
}
