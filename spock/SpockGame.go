package spock

type SpockGame struct {
	strategy Strategy
	lastMove Move
}

func NewSpockGame(strategy Strategy) *SpockGame {
	return &SpockGame{strategy: strategy}
}

func (g *SpockGame) Play(move Move) string {
	other := g.strategy.NextMove(g.lastMove)
	g.lastMove = move

	if verb := move.Verb(other); verb != "" {
		return move.String() + " " + verb + " " + other.String()
	}
	if verb := other.Verb(move); verb != "" {
		return other.String() + " " + verb + " " + move.String()
	}
	return "UNENTSCHIEDEN"
}
