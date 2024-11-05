package spock

type AlwaysStone struct{}

func (s *AlwaysStone) NextMove(lastMove Move) Move {
	return new(SteinMove)
}
