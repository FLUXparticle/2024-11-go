package spock

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockStrategy für das Strategy Interface
type MockStrategy struct {
	mock.Mock
}

func (m *MockStrategy) NextMove(lastMove Move) Move {
	args := m.Called(lastMove)
	return args.Get(0).(Move)
}

func TestSpyMove(t *testing.T) {
	// Nutze die AlwaysStone Strategie ohne Mocking
	game := NewSpockGame(&AlwaysStone{})
	result := game.Play(&PapierMove{})

	assert.Equal(t, "Papier umhüllt Stein", result)
}

func TestMockStrategy(t *testing.T) {
	mockStrategy := new(MockStrategy)
	mockStrategy.On("NextMove", nil).Return(&SteinMove{})

	game := NewSpockGame(mockStrategy)
	result := game.Play(&PapierMove{})

	// Verifiziere das Ergebnis und dass der Mock korrekt aufgerufen wurde
	assert.Equal(t, "Papier umhüllt Stein", result)
	mockStrategy.AssertCalled(t, "NextMove", nil)
}
