package gococktails

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"strings"
	"sync"
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

func getCocktail(id int) *Cocktail {
	response, err := http.Get("http://cocktails.fluxparticle.com/api/cocktails/" + strconv.Itoa(id))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var data *Cocktail // []map[string]any
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data
}

func filterMilkIDs(cocktails []*Cocktail) []int {
	milkIDs := make([]int, 0)
	for _, cocktail := range cocktails {
		if strings.Contains(cocktail.Name, "Milk") {
			milkIDs = append(milkIDs, int(cocktail.CocktailID))
		}
	}

	return milkIDs
}

func TestSequential(t *testing.T) {
	cocktails := getCocktails()

	milkIDs := filterMilkIDs(cocktails)

	sum := 0

	for _, id := range milkIDs {
		cocktail := getCocktail(id)
		//fmt.Printf("%s:\n", cocktail.Name)
		for _, instruction := range cocktail.Instructions {
			if instruction.Ingredient.Name == "Milch" {
				sum += instruction.AmountCL
			}
			//fmt.Printf("  %s\n", instruction.String())
		}
	}

	assert.Equal(t, 38, sum)
}

func TestParallel(t *testing.T) {
	cocktails := getCocktails()

	milkIDs := filterMilkIDs(cocktails)

	ch := make(chan int)

	var wg sync.WaitGroup

	for _, loopID := range milkIDs {
		id := loopID
		wg.Add(1)
		go func() {
			cocktail := getCocktail(id)
			for _, instruction := range cocktail.Instructions {
				if instruction.Ingredient.Name == "Milch" {
					ch <- instruction.AmountCL
				}
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	sum := 0
	for milk := range ch {
		sum += milk
	}

	assert.Equal(t, 38, sum)
}
