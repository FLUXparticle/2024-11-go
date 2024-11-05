package mutex

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"sync"
	"testing"
)

// Funktion zum Zählen der Wörter mit einer normalen Map (ohne Go-Routinen)
func TestWordCount(t *testing.T) {
	file, err := os.Open("story.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			wordCounts[word]++ // Zugriff auf eine nicht-synchronisierte Map
		}
	}

	// Das häufigste Wort bestimmen
	maxWord, maxCount := "", 0
	for word, count := range wordCounts {
		if count > maxCount {
			maxWord, maxCount = word, count
		}
	}

	assert.Equal(t, "ich", maxWord)
	assert.Equal(t, 17, maxCount)
}

// Funktion zum Zählen der Wörter mit einer normalen Map (ohne Synchronisierung)
func TestUnsafeWordCount(t *testing.T) {
	file, err := os.Open("story.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wordCounts := make(map[string]int)
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		// Startet eine Go-Routine pro Zeile
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, word := range words {
				wordCounts[word]++ // Zugriff auf eine nicht-synchronisierte Map
			}
		}()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	wg.Wait()

	// Das häufigste Wort bestimmen
	maxWord, maxCount := "", 0
	for word, count := range wordCounts {
		if count > maxCount {
			maxWord, maxCount = word, count
		}
	}

	assert.Equal(t, "ich", maxWord)
	assert.Equal(t, 17, maxCount)
}

// Funktion zum Zählen der Wörter mit SafeMap (synchronisierte Map)
func TestSafeWordCount(t *testing.T) {
	file, err := os.Open("story.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wordCounts := NewSafeMap()
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		// Startet eine Go-Routine pro Zeile
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, word := range words {
				oldValue, _ := wordCounts.Get(word)
				wordCounts.Set(word, oldValue+1)
			}
		}()
	}

	wg.Wait()

	// Das häufigste Wort bestimmen
	maxWord, maxCount := "", 0
	for word, count := range wordCounts.data {
		if count > maxCount {
			maxWord, maxCount = word, count
		}
	}

	assert.Equal(t, "ich", maxWord)
	assert.Equal(t, 17, maxCount)
}
